package model

/*Context Key*/
type ContextKey string

var Context = struct {
	InstallationIDKey ContextKey
	CustomerIDKey     ContextKey
	LanguageKey       ContextKey
	JWTAct            ContextKey
	FeApiLogNote      ContextKey
}{
	InstallationIDKey: ContextKey("installationID"),
	CustomerIDKey:     ContextKey("customerID"),
	LanguageKey:       ContextKey("language"),
	JWTAct:            ContextKey("act"),
	FeApiLogNote:      ContextKey("apiNote"),
}
