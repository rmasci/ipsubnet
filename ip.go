package ipsubnet

import "strconv"

type Ip struct {
	ip          string
	networkSize int64
	subnet_mask int64
}

func SubnetCalculator(ip string, networkSize int64) *Ip {

	s := &Ip{
		ip:          ip,
		networkSize: networkSize,
		subnet_mask: 0xFFFFFFFF << uint(32-networkSize),
	}

	return s
}

func convertQuardsToInt(splits []string) []int64 {
	quardsInt := []int64{}

	for _, quard := range splits {
		//j, err := strconv.Atoi64(quard)
		j, err := strconv.ParseInt(quard, 10, 64)
		if err != nil {
			panic(err)
		}
		quardsInt = append(quardsInt, j)
	}

	return quardsInt
}
