$params = @{
    "key1" = "value1";
    "key2" = "value2"
}
Invoke-WebRequest -Method POST -Body $params http://localhost:8080/