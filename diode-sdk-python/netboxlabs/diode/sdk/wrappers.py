#!/usr/bin/env python
# Copyright 2024 NetBox Labs Inc
"""NetBox Labs, Diode - SDK - protobuf message wrappers."""

from typing import Mapping as _Mapping
from typing import Optional as _Optional
from typing import Union as _Union

from google.protobuf import timestamp_pb2 as _timestamp_pb2

from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Device as DevicePb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    DeviceType as DeviceTypePb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Entity as EntityPb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Interface as InterfacePb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    IPAddress as IPAddressPb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Manufacturer as ManufacturerPb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Platform as PlatformPb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Prefix as PrefixPb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Role as RolePb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Site as SitePb,
)
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import (
    Tag as TagPb,
)


class Manufacturer:
    """Manufacturer message wrapper."""

    def __new__(
        cls,
        name: _Optional[str] = None,
        slug: _Optional[str] = None,
        description: _Optional[str] = None,
        tags: _Optional[list[str]] = None,
    ) -> ManufacturerPb:
        """Create a new Manufacturer protobuf message."""
        if isinstance(tags, list) and all(isinstance(t, str) for t in tags):
            tags = [TagPb(name=tag) for tag in tags]

        return ManufacturerPb(
            name=name,
            slug=slug,
            description=description,
            tags=tags,
        )


class Platform:
    """Platform message wrapper."""

    def __new__(
        cls,
        name: _Optional[str] = None,
        slug: _Optional[str] = None,
        manufacturer: _Optional[_Union[str, Manufacturer, ManufacturerPb]] = None,
        description: _Optional[str] = None,
        tags: _Optional[list[str]] = None,
    ) -> PlatformPb:
        """Create a new Platform protobuf message."""
        if isinstance(manufacturer, str):
            manufacturer = ManufacturerPb(name=manufacturer)

        if isinstance(tags, list) and all(isinstance(t, str) for t in tags):
            tags = [TagPb(name=tag) for tag in tags]

        return PlatformPb(
            name=name,
            slug=slug,
            manufacturer=manufacturer,
            description=description,
            tags=tags,
        )


class Role:
    """Role message wrapper."""

    def __new__(
        cls,
        name: _Optional[str] = None,
        slug: _Optional[str] = None,
        color: _Optional[str] = None,
        description: _Optional[str] = None,
        tags: _Optional[list[str]] = None,
    ) -> RolePb:
        """Create a new Role protobuf message."""
        if isinstance(tags, list) and all(isinstance(t, str) for t in tags):
            tags = [TagPb(name=tag) for tag in tags]

        return RolePb(
            name=name,
            slug=slug,
            color=color,
            description=description,
            tags=tags,
        )


class DeviceType:
    """DeviceType message wrapper."""

    def __new__(
        cls,
        model: _Optional[str] = None,
        slug: _Optional[str] = None,
        manufacturer: _Optional[_Union[str, Manufacturer, ManufacturerPb]] = None,
        description: _Optional[str] = None,
        comments: _Optional[str] = None,
        part_number: _Optional[str] = None,
        tags: _Optional[list[str]] = None,
    ) -> DeviceTypePb:
        """Create a new DeviceType protobuf message."""
        if isinstance(manufacturer, str):
            manufacturer = ManufacturerPb(name=manufacturer)

        if isinstance(tags, list) and all(isinstance(t, str) for t in tags):
            tags = [TagPb(name=tag) for tag in tags]

        return DeviceTypePb(
            model=model,
            slug=slug,
            manufacturer=manufacturer,
            description=description,
            comments=comments,
            part_number=part_number,
            tags=tags,
        )


class Device:
    """Device message wrapper."""

    def __new__(
        cls,
        name: _Optional[str] = None,
        device_type: _Optional[_Union[str, DeviceType, DeviceTypePb]] = None,
        fqdn: _Optional[str] = None,
        role: _Optional[_Union[str, Role, RolePb]] = None,
        platform: _Optional[_Union[str, Platform, PlatformPb]] = None,
        serial: _Optional[str] = None,
        site: _Optional[_Union[str, SitePb]] = None,
        asset_tag: _Optional[str] = None,
        status: _Optional[str] = None,
        description: _Optional[str] = None,
        comments: _Optional[str] = None,
        tags: _Optional[list[str]] = None,
        primary_ip4: _Optional[_Union[str, IPAddressPb]] = None,
        primary_ip6: _Optional[_Union[str, IPAddressPb]] = None,
        manufacturer: _Optional[_Union[str, Manufacturer, ManufacturerPb]] = None,
    ) -> DevicePb:
        """Create a new Device protobuf message."""
        if isinstance(manufacturer, str):
            manufacturer = ManufacturerPb(name=manufacturer)

        if isinstance(platform, str):
            platform = PlatformPb(name=platform, manufacturer=manufacturer)

        if isinstance(site, str):
            site = SitePb(name=site)

        if isinstance(device_type, str):
            device_type = DeviceTypePb(model=device_type, manufacturer=manufacturer)

        if isinstance(role, str):
            role = RolePb(name=role)

        if isinstance(tags, list) and all(isinstance(t, str) for t in tags):
            tags = [TagPb(name=tag) for tag in tags]

        if isinstance(primary_ip4, str):
            primary_ip4 = IPAddressPb(address=primary_ip4)

        if isinstance(primary_ip6, str):
            primary_ip6 = IPAddressPb(address=primary_ip6)

        return DevicePb(
            name=name,
            device_fqdn=fqdn,
            device_type=device_type,
            role=role,
            platform=platform,
            serial=serial,
            site=site,
            asset_tag=asset_tag,
            status=status,
            description=description,
            comments=comments,
            primary_ip4=primary_ip4,
            primary_ip6=primary_ip6,
            tags=tags,
        )


class Interface:
    """Interface message wrapper."""

    def __new__(
        cls,
        name: _Optional[str] = None,
        device: _Optional[_Union[str, Device, DevicePb]] = None,
        type: _Optional[str] = None,
        enabled: _Optional[bool] = None,
        mtu: _Optional[int] = None,
        mac_address: _Optional[str] = None,
        speed: _Optional[int] = None,
        wwn: _Optional[str] = None,
        mgmt_only: _Optional[bool] = None,
        description: _Optional[str] = None,
        mark_connected: _Optional[bool] = None,
        mode: _Optional[str] = None,
        tags: _Optional[list[str]] = None,
    ) -> InterfacePb:
        """Create a new Interface protobuf message."""
        if isinstance(device, str):
            device = DevicePb(name=device)

        if isinstance(tags, list) and all(isinstance(t, str) for t in tags):
            tags = [TagPb(name=tag) for tag in tags]

        return InterfacePb(
            name=name,
            device=device,
            type=type,
            enabled=enabled,
            mtu=mtu,
            mac_address=mac_address,
            speed=speed,
            wwn=wwn,
            mgmt_only=mgmt_only,
            description=description,
            mark_connected=mark_connected,
            mode=mode,
            tags=tags,
        )


class IPAddress:
    """IPAddress message wrapper."""

    def __new__(
        cls,
        address: _Optional[str] = None,
        interface: _Optional[_Union[str, Interface, InterfacePb]] = None,
        status: _Optional[str] = None,
        role: _Optional[str] = None,
        dns_name: _Optional[str] = None,
        description: _Optional[str] = None,
        comments: _Optional[str] = None,
        tags: _Optional[list[str]] = None,
    ) -> IPAddressPb:
        """Create a new IPAddress protobuf message."""
        if isinstance(interface, str):
            interface = InterfacePb(name=interface)

        if isinstance(tags, list) and all(isinstance(t, str) for t in tags):
            tags = [TagPb(name=tag) for tag in tags]

        return IPAddressPb(
            address=address,
            interface=interface,
            status=status,
            role=role,
            dns_name=dns_name,
            description=description,
            comments=comments,
            tags=tags,
        )


class Prefix:
    """Prefix message wrapper."""

    def __new__(
        cls,
        prefix: _Optional[str] = None,
        site: _Optional[_Union[str, SitePb]] = None,
        status: _Optional[str] = None,
        is_pool: _Optional[bool] = None,
        mark_utilized: _Optional[bool] = None,
        description: _Optional[str] = None,
        comments: _Optional[str] = None,
        tags: _Optional[list[str]] = None,
    ) -> PrefixPb:
        """Create a new Prefix protobuf message."""
        if isinstance(site, str):
            site = SitePb(name=site)

        if isinstance(tags, list) and all(isinstance(t, str) for t in tags):
            tags = [TagPb(name=tag) for tag in tags]

        return PrefixPb(
            prefix=prefix,
            site=site,
            status=status,
            is_pool=is_pool,
            mark_utilized=mark_utilized,
            description=description,
            comments=comments,
            tags=tags,
        )


class Entity:
    """Entity message wrapper."""

    def __new__(
        cls,
        site: _Optional[_Union[str, SitePb]] = None,
        platform: _Optional[_Union[str, Platform, PlatformPb]] = None,
        manufacturer: _Optional[_Union[str, Manufacturer, ManufacturerPb]] = None,
        device: _Optional[_Union[str, Device, DevicePb]] = None,
        device_role: _Optional[_Union[str, Role, RolePb]] = None,
        device_type: _Optional[_Union[str, DeviceType, DeviceTypePb]] = None,
        interface: _Optional[_Union[str, Interface, InterfacePb]] = None,
        ip_address: _Optional[_Union[str, IPAddress, IPAddressPb]] = None,
        prefix: _Optional[_Union[str, Prefix, PrefixPb]] = None,
        timestamp: _Optional[_timestamp_pb2] = None,
    ):
        """Create a new Entity protobuf message."""
        if isinstance(site, str):
            site = SitePb(name=site)

        if isinstance(platform, str):
            platform = PlatformPb(name=platform)

        if isinstance(manufacturer, str):
            manufacturer = ManufacturerPb(name=manufacturer)

        if isinstance(device, str):
            device = DevicePb(name=device)

        if isinstance(device_role, str):
            device_role = RolePb(name=device_role)

        if isinstance(device_type, str):
            device_type = DeviceTypePb(model=device_type)

        if isinstance(ip_address, str):
            ip_address = IPAddressPb(address=ip_address)

        if isinstance(interface, str):
            interface = InterfacePb(name=interface)

        if isinstance(prefix, str):
            prefix = PrefixPb(prefix=prefix)

        return EntityPb(
            site=site,
            platform=platform,
            manufacturer=manufacturer,
            device=device,
            device_role=device_role,
            device_type=device_type,
            interface=interface,
            ip_address=ip_address,
            prefix=prefix,
            timestamp=timestamp,
        )
