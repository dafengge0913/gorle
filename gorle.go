package gorle

const defaultCap = 255
const repeatHead = 1 << 7 // 第一个bit位为1 代表下一个字节是重复字节 后面7个bit是长度
const noRepeatHead = 0
const blockSize = 1<<7 - 1 // 7个bit记录长度

func enBufCap(data []byte) int {
	size := len(data)
	if size < defaultCap {
		return size
	}
	return defaultCap
}

func deBufCap(data []byte) int {
	size := len(data)
	if size > defaultCap {
		return size * 2
	}
	return defaultCap
}

func isRepeat(data []byte, i int) bool {
	if i+2 >= len(data) {
		return false
	}
	return data[i] == data[i+1] && data[i] == data[i+2]
}

func Encode(data []byte) []byte {
	result := make([]byte, 0, enBufCap(data))
	size := len(data)
	index := 0 // 当期处理字节索引
	for index < size {
		count := 1 // 每次至少处理一个字节 data[index]
		if isRepeat(data, index) {
			// 从index之后开始判断是否依然重复
			for i := index + 1; i < size && count < blockSize; i++ {
				if data[i] == data[index] {
					count++
				} else {
					break
				}
			}
			// 记录重复字节长度和重复字节
			result = append(result, repeatHead|byte(count), data[index])
		} else {
			// 从index之后开始判断是否依然不重复
			for i := index + 1; i < size && count < blockSize; i++ {
				if !isRepeat(data, i) {
					count++
				} else {
					break
				}
			}
			// 记录不重复字节长度
			result = append(result, noRepeatHead|byte(count))
			// 复制不重复的那些字节
			result = append(result, data[index:index+count]...)
		}
		index += count
	}
	return result
}

func Decode(data []byte) []byte {
	result := make([]byte, 0, deBufCap(data))
	size := len(data)
	index := 0 // 当期处理字节索引
	for index < size {
		count := 0
		if data[index]&repeatHead == repeatHead {
			// 处理重复字节
			count = int(data[index] ^ repeatHead)
			b := data[index+1] // 下一个字节记录着 被重复的字节
			for i := 0; i < count; i++ {
				result = append(result, b)
			}
			index += 2 // 处理了两个字节 头字节和重复字节
		} else {
			count = int(data[index] ^ noRepeatHead)
			for i := 0; i < count; i++ {
				result = append(result, data[index+i+1])
			}
			index += count + 1 // 处理头字节和count个不重复字节
		}
	}
	return result
}
