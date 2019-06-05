#!/usr/bin/env python3
import pychromecast
import time
import grpc
from uuid import UUID
from concurrent import futures
import cast_control_pb2
import cast_control_pb2_grpc
import logging
from grpc_reflection.v1alpha import reflection

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

class CastControlServer(cast_control_pb2_grpc.CastControlServicer):
    def __init__(self):
        self.chromecasts = pychromecast.get_chromecasts()
    
    def ListDevices(self, request, context):
        print(f"There are {len(self.chromecasts)} chromecasts")

        def grpc_map(chromecast: pychromecast.Chromecast) -> cast_control_pb2.Device:
            return cast_control_pb2.Device(
                friendly_name = chromecast.device.friendly_name,
                model_name = chromecast.device.model_name,
                manufacturer = chromecast.device.manufacturer,
                device_id = cast_control_pb2.DeviceID(uuid=str(chromecast.device.uuid)),
                cast_type = chromecast.device.cast_type
            )

        for cc in map(grpc_map, self.chromecasts):
            yield cc

    def GetDeviceStatus(self, request: cast_control_pb2.DeviceID, context: grpc.ServicerContext):
        print(f"Got request for UUID '{request.uuid}'")
        
        # Check if the string is actually a UUID
        try:
            uuid = UUID(request.uuid)
        except ValueError:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details(f"Illegal UUID: '{request.uuid}'")
            return cast_control_pb2.Status()

        # Check if the ID actually exists on the network
        chromecast = next((c for c in self.chromecasts if c.device.uuid == uuid), None)
        if chromecast == None:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details(f"No device with UUID '{request.uuid}' found on the network")
            return cast_control_pb2.Status()

        chromecast.wait()
        cc_status = chromecast.status
        return cast_control_pb2.Status(volume=cc_status.volume_level, muted=cc_status.volume_muted)

    def AdjustVolume(self, request: cast_control_pb2.AdjustVolumeRequest, context):
        print(f"Got request for UUID '{request.device_id.uuid}'")
        
        # Check if the string is actually a UUID
        try:
            uuid = UUID(request.device_id.uuid)
        except ValueError:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details(f"Illegal UUID: '{request.device_id.uuid}'")
            return cast_control_pb2.Status()

        # Check if the ID actually exists on the network
        chromecast = next((c for c in self.chromecasts if c.device.uuid == uuid), None)
        if chromecast == None:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details(f"No device with UUID '{request.device_id.uuid}' found on the network")
            return cast_control_pb2.Volume()

        # Adjusting the volume by 0% is a no-op
        if request.relative_volume == 0:
            context.set_code(grpc.StatusCode.OUT_OF_RANGE)
            context.set_details("Volume cannot be adjusted by 0%")
            return cast_control_pb2.Volume()
        
        chromecast.wait()
        new_volume = None
        if request.relative_volume > 0:
            new_volume = chromecast.volume_up(delta=request.relative_volume)
        elif request.relative_volume < 0:
            new_volume = chromecast.volume_down(delta=request.relative_volume * -1)
        
        return cast_control_pb2.Volume(volume=new_volume)

    def SetVolume(self, request: cast_control_pb2.SetVolumeRequest, context):
        print(f"Got request for UUID '{request.device_id.uuid}'")
        
        # Check if the string is actually a UUID
        try:
            uuid = UUID(request.device_id.uuid)
        except ValueError:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details(f"Illegal UUID: '{request.device_id.uuid}'")
            return cast_control_pb2.Status()

        # Check if the ID actually exists on the network
        chromecast = next((c for c in self.chromecasts if c.device.uuid == uuid), None)
        if chromecast == None:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details(f"No device with UUID '{request.device_id.uuid}' found on the network")
            return cast_control_pb2.Volume()
        
        chromecast.wait()
        new_volume = chromecast.set_volume(volume=request.volume)
        return cast_control_pb2.Volume(volume=new_volume)

def serve(listen_addr: str) -> None:
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cast_control_pb2_grpc.add_CastControlServicer_to_server(
        CastControlServer(), server)
    SERVICE_NAMES = (cast_control_pb2.DESCRIPTOR.services_by_name['CastControl'].full_name, reflection.SERVICE_NAME, )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    server.add_insecure_port(listen_addr)
    server.start()
    print(f"Server started, listening on {listen_addr}")
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == "__main__":
    addr = '[::]:7004'
    logging.basicConfig()
    print("Bootstrapping, this may take a minute...")
    serve(addr)
