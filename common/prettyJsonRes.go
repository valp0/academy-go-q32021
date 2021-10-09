package common

func PrettyJsonRes(r interface{}) string {
	return PrettifyJson(JsonResponse(r))
}
