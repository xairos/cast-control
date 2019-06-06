package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/OpenPeeDeeP/xdg"
	pb "github.com/xairos/cast-control/client/protocol"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app   = kingpin.New("cast-control", "Control Chromecast devices on your network.")
	debug = kingpin.Flag("debug", "Enable debug mode.").Bool()

	configureCmd     = app.Command("configure", "Interactively set the tool's configuration values.")
	volumeCmd        = app.Command("volume", "Control the volume.")
	volumeUpCmd      = volumeCmd.Command("up", "Increase the volume.")
	volumeUpAmount   = volumeUpCmd.Arg("amount", "Amount to increase the volume by (0-1).").Required().Float()
	volumeDownCmd    = volumeCmd.Command("down", "Decrease the volume.")
	volumeDownAmount = volumeDownCmd.Arg("amount", "Amount to decrease the volume by (0-1).").Required().Float()
	setVolumeCmd     = volumeCmd.Command("set", "Set the volume.")
	setVolumeLevel   = setVolumeCmd.Arg("level", "Level to set the volume to (0-1).").Required().Float()
	statusCmd        = app.Command("status", "Get the status of the device.")
)

var defaultServerAddress = "localhost:7004"

// Cached configuration format
type config struct {
	ServerAddress  string `json:"serverAddress"`
	ChromecastUUID string `json:"chromecastUUID"`
}

// GRPCChromecastController wraps the cast-control server to control a single device
type GRPCChromecastController struct {
	client *pb.CastControlClient
	uuid   string
}

func (c *GRPCChromecastController) adjustVolume(amount float64) float64 {
	// TODO: Error-check for a 0-level adjustment, which is invalid
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	newVolume, err := (*c.client).AdjustVolume(
		ctx,
		&pb.AdjustVolumeRequest{
			DeviceId:       &pb.DeviceID{Uuid: c.uuid},
			RelativeVolume: amount})
	if err != nil {
		log.Fatalf("%v.AdjustVolume(_) = _, %v", *c.client, err)
	}
	return newVolume.GetVolume()
}

func (c *GRPCChromecastController) setVolume(volumeLevel float64) float64 {
	// TODO: Error check acceptable range (0-1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	newVolume, err := (*c.client).SetVolume(
		ctx,
		&pb.SetVolumeRequest{
			DeviceId: &pb.DeviceID{Uuid: c.uuid},
			Volume:   volumeLevel})
	if err != nil {
		log.Fatalf("%v.SetVolume(_) = _, %v", *c.client, err)
	}
	return newVolume.GetVolume()
}

func (c *GRPCChromecastController) getStatus() *pb.Status {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	status, err := (*c.client).GetDeviceStatus(ctx, &pb.DeviceID{Uuid: c.uuid})
	if err != nil {
		log.Fatalf("%v.GetDeviceStatus(_) = _, %v", *c.client, err)
	}
	return status
}

func listDevices(client pb.CastControlClient) []*pb.Device {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := client.ListDevices(ctx, &pb.ListDeviceRequest{})
	if err != nil {
		log.Fatalf("%v.ListDevices(_) = _, %v: ", client, err)
	}

	devices := []*pb.Device{}
	for {
		device, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListDevices(_) = _, %v", client, err)
		}
		devices = append(devices, device)
	}
	return devices
}

/* Workflow
Server address (ex. localhost:7004)
Trying to connect...
Failed (loop) / Success.
? Choose a device:
  TV - P55-F1
  Bedroom Mini - Google Home Mini
> Bedroom Speakers - Chromecast Audio
  Bedroom speaker - Google Home Mini
*/
func configureWizard() *config {
	serverAddress := ""
	var conn *grpc.ClientConn
	for {
		serverAddressPrompt := &survey.Input{
			Message: "Server address (ex. localhost:7004)",
		}
		survey.AskOne(serverAddressPrompt, &serverAddress, survey.WithValidator(survey.Required))

		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		conn, err = grpc.DialContext(ctx, serverAddress, grpc.WithBlock(), grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Failed to connect to address: %s\n", serverAddress)
			continue
		}
		break
	}
	defer conn.Close()

	client := pb.NewCastControlClient(conn)
	chromecasts := listDevices(client)

	// Build a map from option -> device
	deviceMap := make(map[string]*pb.Device, len(chromecasts))
	keys := make([]string, len(chromecasts))
	for i, chromecast := range chromecasts {
		// User selects from a list of "FriendlyName - ModelName"
		key := fmt.Sprintf("%s - %s", chromecast.GetFriendlyName(), chromecast.GetModelName())
		keys[i] = key
		deviceMap[key] = chromecast
	}

	deviceSelection := ""
	selectorPrompt := &survey.Select{
		Message: "Select a device:",
		Options: keys,
	}
	survey.AskOne(selectorPrompt, &deviceSelection)

	return &config{
		ServerAddress:  serverAddress,
		ChromecastUUID: deviceMap[deviceSelection].GetDeviceId().Uuid}
}

func main() {
	configFolder := xdg.New("", "cast-control").ConfigHome()
	configFilePath := path.Join(configFolder, "config.json")
	conf := &config{}

	// Version must be declared before parsing is performed
	app.Version("1.0.0")
	commandString := kingpin.MustParse(app.Parse(os.Args[1:]))

	// Try commands that don't require a connection to the server first
	cmdHandled := true
	switch commandString {
	case "configure":
		// Launch the configuration wizard to prompt the user for values
		configToWrite := configureWizard()

		// Create the config folder if it doesn't exist already
		if _, err := os.Stat(configFolder); os.IsNotExist(err) {
			err = os.Mkdir(configFolder, 0700)
			if err != nil {
				log.Fatalf("Couldn't create configuration path: %v", err)
			}
		}

		// Write out the JSON configuration file
		jsonData, err := json.Marshal(configToWrite)
		if err != nil {
			log.Fatalf("Couldn't marshal config to json: %v", err)
		}
		err = ioutil.WriteFile(configFilePath, jsonData, 0644)
		if err != nil {
			log.Fatalf("Couldn't write out configuration file: %v", err)
		}
		fmt.Println("Configuration saved.")

	default:
		cmdHandled = false
	}
	if cmdHandled {
		return
	}

	// If the configuration file doesn't exist, this is probably a first-time run
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		fmt.Println("It looks like this is the first time running cast-control.")
		fmt.Println("Run the `configure` command to set up a device.")
		return
	}

	// Read the configuration file into memory
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Could not open configuration file: %v", err)
	}
	err = json.Unmarshal(content, conf)
	if err != nil {
		log.Fatalf("Invalid config file format: %v", err)
	}

	// Open a connection and client to the server in the configuration
	conn, err := grpc.Dial(conf.ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCastControlClient(conn)
	controller := GRPCChromecastController{client: &client, uuid: conf.ChromecastUUID}

	switch commandString {
	case "volume up":
		newVolume := controller.adjustVolume(*volumeUpAmount)
		fmt.Printf("Volume increased to %.3f\n", newVolume)

	case "volume down":
		newVolume := controller.adjustVolume(*volumeDownAmount * -1)
		fmt.Printf("Decreased volume to %.3f\n", newVolume)

	case "volume set":
		newVolume := controller.setVolume(*setVolumeLevel)
		fmt.Printf("Volume set to %.3f\n", newVolume)

	case "status":
		status := controller.getStatus()
		fmt.Printf("Volume=%.3f, Muted=%t\n", status.GetVolume(), status.GetMuted())
	}
}
