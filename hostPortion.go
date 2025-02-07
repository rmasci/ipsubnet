package ipsubnet

import (
	"fmt"
	"strings"
)

func (s *Ip) GetHostPortion() string {
	return s.hostCalculation("%d", ".")
}

func (s *Ip) GetHostPortionQuards() []int64 {
	return convertQuardsToInt(strings.Split(s.hostCalculation("%d", "."), "."))
}

func (s *Ip) GetHostPortionHex() string {
	return s.hostCalculation("%02X", "")
}

func (s *Ip) GetHostPortionBinary() string {
	return s.hostCalculation("%08b", "")
}

func (s *Ip) hostCalculation(format, separator string) string {
	splits := s.GetIPAddressQuads()
	networkQuards := []string{}
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[0] & ^(s.subnet_mask>>int64(24))))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[1] & ^(s.subnet_mask>>int64(16))))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[2] & ^(s.subnet_mask>>int64(8))))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[3] & ^(s.subnet_mask>>int64(0))))

	return strings.Join(networkQuards, separator)
}
