package dp

import (
	"encoding/binary"
	"strconv"
	"strings"
	"testing"
)

func TestUnPack(t *testing.T) {
	var (
		txt = `0 = 143
1 = 0
2 = 0
3 = 0
4 = 1
5 = 123
6 = 34
7 = 116
8 = 114
9 = 97
10 = 99
11 = 101
12 = 73
13 = 100
14 = 34
15 = 58
16 = 34
17 = 54
18 = 53
19 = 97
20 = 99
21 = 55
22 = 52
23 = 54
24 = 102
25 = 50
26 = 100
27 = 101
28 = 101
29 = 49
30 = 52
31 = 57
32 = 98
33 = 54
34 = 97
35 = 51
36 = 51
37 = 100
38 = 98
39 = 54
40 = 100
41 = 102
42 = 51
43 = 54
44 = 102
45 = 98
46 = 48
47 = 56
48 = 49
49 = 34
50 = 44
51 = 34
52 = 99
53 = 111
54 = 100
55 = 101
56 = 34
57 = 58
58 = 48
59 = 44
60 = 34
61 = 100
62 = 97
63 = 116
64 = 97
65 = 34
66 = 58
67 = 34
68 = 123
69 = 92
70 = 34
71 = 100
72 = 105
73 = 114
74 = 92
75 = 34
76 = 58
77 = 92
78 = 34
79 = 47
80 = 100
81 = 97
82 = 116
83 = 97
84 = 47
85 = 117
86 = 115
87 = 101
88 = 114
89 = 47
90 = 48
91 = 47
92 = 111
93 = 114
94 = 103
95 = 46
96 = 102
97 = 114
98 = 101
99 = 101
100 = 100
101 = 105
102 = 109
103 = 46
104 = 102
105 = 100
106 = 108
107 = 105
108 = 98
109 = 46
110 = 102
111 = 100
112 = 108
113 = 105
114 = 98
115 = 95
116 = 102
117 = 108
118 = 117
119 = 116
120 = 116
121 = 101
122 = 114
123 = 95
124 = 101
125 = 120
126 = 97
127 = 109
128 = 112
129 = 108
130 = 101
131 = 47
132 = 97
133 = 112
134 = 112
135 = 95
136 = 102
137 = 108
138 = 117
139 = 116
140 = 116
141 = 101
142 = 114
143 = 92
144 = 34
145 = 125
146 = 34
147 = 125`
		data = make([]byte, 0)
	)
	// 解析txt
	lines := strings.Split(txt, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		kv := strings.Split(line, "=")
		if len(kv) != 2 {
			continue
		}
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		if k == "" || v == "" {
			continue
		}
		vi, _ := strconv.ParseInt(v, 10, 64)
		data = append(data, byte(vi))
	}
	// 解析data
	//先读出dataLen
	headData := make([]byte, 5)
	copy(headData, data[:5])
	//只解压head的信息，得到dataLen和msgId
	msg := &Message{
		Len: binary.LittleEndian.Uint32(headData[:4]),
	}
	//读isResponse
	msg.IsResponse = headData[4] == 1
	dataTmp := make([]byte, msg.Len)
	//读data数据
	copy(dataTmp, data[5:])
	msg.Data = string(dataTmp)
	t.Logf("msg: %+v\n", msg)
}
