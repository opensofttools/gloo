package utils

import (
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
)

func PathAsString(matcher *matchers.Matcher) string {
	switch path := matcher.GetPathSpecifier().(type) {
	case *matchers.Matcher_Prefix:
		return path.Prefix
	case *matchers.Matcher_Exact:
		return path.Exact
	case *matchers.Matcher_Regex:
		return path.Regex
	}
	return ""
}

func EnvoyPathAsString(matcher *route.RouteMatch) string {
	switch path := matcher.GetPathSpecifier().(type) {
	//	TODO add the other cases? Didn't have SafeRegex before either
	//	*RouteMatch_SafeRegex
	//	*RouteMatch_ConnectMatcher_
	case *route.RouteMatch_Prefix:
		return path.Prefix
	case *route.RouteMatch_Path:
		return path.Path
	case *route.RouteMatch_HiddenEnvoyDeprecatedRegex:
		return path.HiddenEnvoyDeprecatedRegex
	}
	return ""
}
