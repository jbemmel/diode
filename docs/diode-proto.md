# Protocol Documentation

<a name="top"></a>

## Table of Contents

- [diode/v1/ingester.proto](#diode_v1_ingester-proto)
    - [Device](#diode-v1-Device)
    - [DeviceType](#diode-v1-DeviceType)
    - [Entity](#diode-v1-Entity)
    - [IPAddress](#diode-v1-IPAddress)
    - [IngestRequest](#diode-v1-IngestRequest)
    - [IngestResponse](#diode-v1-IngestResponse)
    - [Interface](#diode-v1-Interface)
    - [Manufacturer](#diode-v1-Manufacturer)
    - [Platform](#diode-v1-Platform)
    - [Prefix](#diode-v1-Prefix)
    - [Role](#diode-v1-Role)
    - [Site](#diode-v1-Site)
    - [Tag](#diode-v1-Tag)

    - [IngesterService](#diode-v1-IngesterService)

- [Scalar Value Types](#scalar-value-types)

<a name="diode_v1_ingester-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## diode/v1/ingester.proto

<a name="diode-v1-Device"></a>

### Device

A device

| Field       | Type                               | Label    | Description |
|-------------|------------------------------------|----------|-------------|
| name        | [string](#string)                  |          |             |
| device_fqdn | [string](#string)                  |          |             |
| device_type | [DeviceType](#diode-v1-DeviceType) |          |             |
| role        | [Role](#diode-v1-Role)             |          |             |
| platform    | [Platform](#diode-v1-Platform)     |          |             |
| serial      | [string](#string)                  |          |             |
| site        | [Site](#diode-v1-Site)             |          |             |
| asset_tag   | [string](#string)                  |          |             |
| status      | [string](#string)                  |          |             |
| description | [string](#string)                  |          |             |
| comments    | [string](#string)                  |          |             |
| tags        | [Tag](#diode-v1-Tag)               | repeated |             |
| primary_ip4 | [IPAddress](#diode-v1-IPAddress)   |          |             |
| primary_ip6 | [IPAddress](#diode-v1-IPAddress)   |          |             |

<a name="diode-v1-DeviceType"></a>

### DeviceType

A device type

| Field        | Type                                   | Label    | Description |
|--------------|----------------------------------------|----------|-------------|
| model        | [string](#string)                      |          |             |
| slug         | [string](#string)                      |          |             |
| manufacturer | [Manufacturer](#diode-v1-Manufacturer) |          |             |
| description  | [string](#string)                      |          |             |
| comments     | [string](#string)                      |          |             |
| part_number  | [string](#string)                      |          |             |
| tags         | [Tag](#diode-v1-Tag)                   | repeated |             |

<a name="diode-v1-Entity"></a>

### Entity

An ingest entity wrapper

| Field        | Type                                                    | Label | Description                                   |
|--------------|---------------------------------------------------------|-------|-----------------------------------------------|
| site         | [Site](#diode-v1-Site)                                  |       |                                               |
| platform     | [Platform](#diode-v1-Platform)                          |       |                                               |
| manufacturer | [Manufacturer](#diode-v1-Manufacturer)                  |       |                                               |
| device       | [Device](#diode-v1-Device)                              |       |                                               |
| device_role  | [Role](#diode-v1-Role)                                  |       |                                               |
| device_type  | [DeviceType](#diode-v1-DeviceType)                      |       |                                               |
| interface    | [Interface](#diode-v1-Interface)                        |       |                                               |
| ip_address   | [IPAddress](#diode-v1-IPAddress)                        |       |                                               |
| prefix       | [Prefix](#diode-v1-Prefix)                              |       |                                               |
| timestamp    | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |       | The timestamp of the data discovery at source |

<a name="diode-v1-IPAddress"></a>

### IPAddress

An IP address.

| Field       | Type                             | Label    | Description |
|-------------|----------------------------------|----------|-------------|
| address     | [string](#string)                |          |             |
| interface   | [Interface](#diode-v1-Interface) |          |             |
| status      | [string](#string)                |          |             |
| role        | [string](#string)                |          |             |
| dns_name    | [string](#string)                |          |             |
| description | [string](#string)                |          |             |
| comments    | [string](#string)                |          |             |
| tags        | [Tag](#diode-v1-Tag)             | repeated |             |

<a name="diode-v1-IngestRequest"></a>

### IngestRequest

The request to ingest the data

| Field                | Type                       | Label    | Description |
|----------------------|----------------------------|----------|-------------|
| stream               | [string](#string)          |          |             |
| entities             | [Entity](#diode-v1-Entity) | repeated |             |
| id                   | [string](#string)          |          |             |
| producer_app_name    | [string](#string)          |          |             |
| producer_app_version | [string](#string)          |          |             |
| sdk_name             | [string](#string)          |          |             |
| sdk_version          | [string](#string)          |          |             |

<a name="diode-v1-IngestResponse"></a>

### IngestResponse

The response from the ingest request

| Field  | Type              | Label    | Description |
|--------|-------------------|----------|-------------|
| errors | [string](#string) | repeated |             |

<a name="diode-v1-Interface"></a>

### Interface

An interface

| Field          | Type                       | Label    | Description |
|----------------|----------------------------|----------|-------------|
| device         | [Device](#diode-v1-Device) |          |             |
| name           | [string](#string)          |          |             |
| label          | [string](#string)          |          |             |
| type           | [string](#string)          |          |             |
| enabled        | [bool](#bool)              |          |             |
| mtu            | [int32](#int32)            |          |             |
| mac_address    | [string](#string)          |          |             |
| speed          | [int32](#int32)            |          |             |
| wwn            | [string](#string)          |          |             |
| mgmt_only      | [bool](#bool)              |          |             |
| description    | [string](#string)          |          |             |
| mark_connected | [bool](#bool)              |          |             |
| mode           | [string](#string)          |          |             |
| tags           | [Tag](#diode-v1-Tag)       | repeated |             |

<a name="diode-v1-Manufacturer"></a>

### Manufacturer

A manufacturer

| Field       | Type                 | Label    | Description |
|-------------|----------------------|----------|-------------|
| name        | [string](#string)    |          |             |
| slug        | [string](#string)    |          |             |
| description | [string](#string)    |          |             |
| tags        | [Tag](#diode-v1-Tag) | repeated |             |

<a name="diode-v1-Platform"></a>

### Platform

A platform

| Field        | Type                                   | Label    | Description |
|--------------|----------------------------------------|----------|-------------|
| name         | [string](#string)                      |          |             |
| slug         | [string](#string)                      |          |             |
| manufacturer | [Manufacturer](#diode-v1-Manufacturer) |          |             |
| description  | [string](#string)                      |          |             |
| tags         | [Tag](#diode-v1-Tag)                   | repeated |             |

<a name="diode-v1-Prefix"></a>

### Prefix

An IPAM prefix.

| Field         | Type                   | Label    | Description |
|---------------|------------------------|----------|-------------|
| prefix        | [string](#string)      |          |             |
| site          | [Site](#diode-v1-Site) |          |             |
| status        | [string](#string)      |          |             |
| is_pool       | [bool](#bool)          |          |             |
| mark_utilized | [bool](#bool)          |          |             |
| description   | [string](#string)      |          |             |
| comments      | [string](#string)      |          |             |
| tags          | [Tag](#diode-v1-Tag)   | repeated |             |

<a name="diode-v1-Role"></a>

### Role

A role

| Field       | Type                 | Label    | Description |
|-------------|----------------------|----------|-------------|
| name        | [string](#string)    |          |             |
| slug        | [string](#string)    |          |             |
| color       | [string](#string)    |          |             |
| description | [string](#string)    |          |             |
| tags        | [Tag](#diode-v1-Tag) | repeated |             |

<a name="diode-v1-Site"></a>

### Site

A site

| Field       | Type                 | Label    | Description |
|-------------|----------------------|----------|-------------|
| name        | [string](#string)    |          |             |
| slug        | [string](#string)    |          |             |
| status      | [string](#string)    |          |             |
| facility    | [string](#string)    |          |             |
| time_zone   | [string](#string)    |          |             |
| description | [string](#string)    |          |             |
| comments    | [string](#string)    |          |             |
| tags        | [Tag](#diode-v1-Tag) | repeated |             |

<a name="diode-v1-Tag"></a>

### Tag

A tag

| Field | Type              | Label | Description |
|-------|-------------------|-------|-------------|
| name  | [string](#string) |       |             |
| slug  | [string](#string) |       |             |
| color | [string](#string) |       |             |

<a name="diode-v1-IngesterService"></a>

### IngesterService

Ingestion API

| Method Name | Request Type                             | Response Type                              | Description                  |
|-------------|------------------------------------------|--------------------------------------------|------------------------------|
| Ingest      | [IngestRequest](#diode-v1-IngestRequest) | [IngestResponse](#diode-v1-IngestResponse) | Ingests data into the system |

## Scalar Value Types

| .proto Type                    | Notes                                                                                                                                           | C++    | Java       | Python      | Go      | C#         | PHP            | Ruby                           |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|--------|------------|-------------|---------|------------|----------------|--------------------------------|
| <a name="double" /> double     |                                                                                                                                                 | double | double     | float       | float64 | double     | float          | Float                          |
| <a name="float" /> float       |                                                                                                                                                 | float  | float      | float       | float32 | float      | float          | Float                          |
| <a name="int32" /> int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="int64" /> int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="uint32" /> uint32     | Uses variable-length encoding.                                                                                                                  | uint32 | int        | int/long    | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64     | Uses variable-length encoding.                                                                                                                  | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="fixed32" /> fixed32   | Always four bytes. More efficient than uint32 if values are often greater than 2^28.                                                            | uint32 | int        | int         | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64   | Always eight bytes. More efficient than uint64 if values are often greater than 2^56.                                                           | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum                         |
| <a name="sfixed32" /> sfixed32 | Always four bytes.                                                                                                                              | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes.                                                                                                                             | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="bool" /> bool         |                                                                                                                                                 | bool   | boolean    | boolean     | bool    | bool       | boolean        | TrueClass/FalseClass           |
| <a name="string" /> string     | A string must always contain UTF-8 encoded or 7-bit ASCII text.                                                                                 | string | String     | str/unicode | string  | string     | string         | String (UTF-8)                 |
| <a name="bytes" /> bytes       | May contain any arbitrary sequence of bytes.                                                                                                    | string | ByteString | str         | []byte  | ByteString | string         | String (ASCII-8BIT)            |

