package runtime

import "github.com/opencontainers/specs"

func getRootIDs(s *platformSpec) (int, int, error) {
	if s == nil {
		return 0, 0, nil
	}
	var hasUserns bool
	for _, ns := range s.Spec.Linux.Namespaces {
		if ns.Type == specs.UserNamespace {
			hasUserns = true
			break
		}
	}
	if !hasUserns {
		return 0, 0, nil
	}
	uid := hostIDFromMap(0, s.Spec.Linux.UIDMappings)
	gid := hostIDFromMap(0, s.Spec.Linux.GIDMappings)
	return uid, gid, nil
}
