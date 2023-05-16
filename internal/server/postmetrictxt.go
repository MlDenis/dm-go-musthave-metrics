package server

vt := ctx.Params.ByName(metrics.TypeS)
name := ctx.Params.ByName(metrics.NameS)
value := ctx.Params.ByName(metrics.ValueS)

log.Printf("#DEBUG run PostSingleValue with: value type = %s, name = %s, value = %s.\n", vt, name, value)

if value == "none" {
ctx.String(http.StatusBadRequest, services.CtText)
return
}

switch vt {
case metrics.GaugeS:
if _, err := strconv.ParseFloat(value, 64); err == nil {
s.MetricsStorage.UpdateStorageData(vt, name, value)
ctx.String(http.StatusOK, services.CtText)
return
}
case metrics.CounterS:
log.Printf("#DEBUG.%s.%s value before strconv.ParseInt(value, 10, 64) = %s", vt, name, value)
if i, err := strconv.ParseInt(value, 10, 64); err == nil {
log.Printf("#DEBUG.%s.%s value before strconv.ParseInt(value, 10, 64) = %v", vt, name, i)
s.MetricsStorage.UpdateStorageData(vt, name, value)
ctx.String(http.StatusOK, services.CtText)
return
}
default:
ctx.String(http.StatusNotImplemented, services.CtText)
return
}
