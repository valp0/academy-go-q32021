package common

// Will return a prettified JSON string containing what it receives.
// Equivalent to ``PrettifyJson(JsonResponse(r))``
func PrettyJsonRes(r interface{}) string {
	return PrettifyJson(JsonResponse(r))
}
