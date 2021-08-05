# Seller Authentication Services

[![N|Solid](https://ban.sendo.vn/Content/images/design/logo/logo-sendo-ban-161x32.png)](https://ban.sendo.vn/)

Seller Authentication Services cung cấp các api login shop, lấy danh sách menu của shop, phân quyền shop, nhân viên được phân quyền.
(https://sendovn.atlassian.net/wiki/spaces/PRC/pages/1552187613/PSD+New+Sign+up+Sign+up+of+Seller)


### Tech

Seller Authentication Services uses a number of open source projects to work properly:

* [.NET CORE 5]
* [Dapper C#/.SQL Driver]
* [StackExchange.Redis]
* [EasyNetQ]



### Installation

Seller Authentication Services requires [.NET V5](https://dotnet.microsoft.com/download/dotnet/5.0) 

Install the dependencies and devDependencies and start the server.

```sh
$ cd src\Runnable\Sendo.Seller.Shop.Authentication.Api
$ dotnet run --environment=Production
```

For production environments...

```sh
$ dotnet publish "src/Runnable/Sendo.Seller.Shop.Authentication.Api/Sendo.Seller.Shop.Authentication.Api.csproj" --configuration Release --output _output/Sendo.Seller.Shop.Authentication.Api

$ dotnet Sendo.Seller.Shop.Authentication.Api.dll --environment=Production --urls "http://localhost:5001"
```

### Development

Want to contribute? Great!
Contact nguyenpt8@sendo.vn


