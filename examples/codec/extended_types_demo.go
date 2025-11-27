package main

import (
	"bytes"
	"fmt"

	"github.com/danclive/nson-go"
)

func main() {
	fmt.Println("=== NSON Extended Types Demo ===")

	// 示例 1: 基本类型
	fmt.Println("1. Basic Types:")
	demonstrateBasicTypes()

	// 示例 2: Matter 设备信息
	fmt.Println("\n2. Matter Device Info:")
	demonstrateMatterDevice()

	// 示例 3: OnOff Cluster
	fmt.Println("\n3. OnOff Cluster State:")
	demonstrateOnOffCluster()

	// 示例 4: 温度传感器
	fmt.Println("\n4. Temperature Sensor:")
	demonstrateTemperatureSensor()

	// 示例 5: 序列化和反序列化
	fmt.Println("\n5. Serialization:")
	demonstrateSerialization()
}

func demonstrateBasicTypes() {
	values := map[string]nson.Value{
		"U8":  nson.U8(255),
		"U16": nson.U16(65535),
		"I8":  nson.I8(-128),
		"I16": nson.I16(-32768),
	}

	for name, val := range values {
		fmt.Printf("  %-4s: %v\n", name, val)
	}
}

func demonstrateMatterDevice() {
	// 创建 Matter 设备基本信息
	deviceInfo := nson.Map{
		"vendorId":        nson.U16(0x1234),
		"productId":       nson.U16(0x5678),
		"vendorName":      nson.String("Acme Corp"),
		"productName":     nson.String("Smart Light Pro"),
		"serialNumber":    nson.String("SN-2024-001"),
		"hardwareVersion": nson.U16(2),
		"softwareVersion": nson.U32(0x00010203), // v1.2.3
	}

	// 显示设备信息
	if vendorId, ok := deviceInfo["vendorId"].(nson.U16); ok {
		fmt.Printf("  Vendor ID: 0x%04X\n", vendorId)
	}
	if vendorName, ok := deviceInfo["vendorName"].(nson.String); ok {
		fmt.Printf("  Vendor: %s\n", vendorName)
	}
	if productName, ok := deviceInfo["productName"].(nson.String); ok {
		fmt.Printf("  Product: %s\n", productName)
	}
	if swVer, ok := deviceInfo["softwareVersion"].(nson.U32); ok {
		major := (swVer >> 16) & 0xFF
		minor := (swVer >> 8) & 0xFF
		patch := swVer & 0xFF
		fmt.Printf("  Software Version: v%d.%d.%d\n", major, minor, patch)
	}
}

func demonstrateOnOffCluster() {
	// OnOff Cluster 状态
	onOffState := nson.Map{
		"clusterID":          nson.U32(0x0006), // OnOff Cluster ID
		"onOff":              nson.Bool(true),
		"globalSceneControl": nson.Bool(true),
		"onTime":             nson.U16(0),
		"offWaitTime":        nson.U16(0),
	}

	if onOff, ok := onOffState["onOff"].(nson.Bool); ok {
		status := "OFF"
		if onOff {
			status = "ON"
		}
		fmt.Printf("  Current State: %s\n", status)
	}

	// 模拟状态变化
	onOffState["onOff"] = nson.Bool(false)
	fmt.Println("  State changed to: OFF")
}

func demonstrateTemperatureSensor() {
	// 温度传感器数据（单位：0.01°C）
	tempSensor := nson.Map{
		"clusterID":        nson.U32(0x0402), // Temperature Measurement
		"measuredValue":    nson.I16(2350),   // 23.50°C
		"minMeasuredValue": nson.I16(-5000),  // -50.00°C
		"maxMeasuredValue": nson.I16(10000),  // 100.00°C
		"tolerance":        nson.U16(100),    // ±1.00°C
	}

	// 读取并显示温度
	if temp, ok := tempSensor["measuredValue"].(nson.I16); ok {
		actualTemp := float64(temp) / 100.0
		fmt.Printf("  Current Temperature: %.2f°C\n", actualTemp)
	}

	if minTemp, ok := tempSensor["minMeasuredValue"].(nson.I16); ok {
		fmt.Printf("  Min Temperature: %.2f°C\n", float64(minTemp)/100.0)
	}

	if maxTemp, ok := tempSensor["maxMeasuredValue"].(nson.I16); ok {
		fmt.Printf("  Max Temperature: %.2f°C\n", float64(maxTemp)/100.0)
	}

	if tolerance, ok := tempSensor["tolerance"].(nson.U16); ok {
		fmt.Printf("  Tolerance: ±%.2f°C\n", float64(tolerance)/100.0)
	}
}

func demonstrateSerialization() {
	// 创建一个复杂的设备状态
	deviceState := nson.Map{
		"deviceId":    nson.U16(1),
		"name":        nson.String("Living Room Light"),
		"online":      nson.Bool(true),
		"brightness":  nson.U8(200),
		"temperature": nson.I16(2250),
		"endpoints": nson.Array{
			nson.Map{
				"id":   nson.U16(1),
				"type": nson.U32(0x0101), // DimmableLight
			},
		},
	}

	// 序列化
	var buf bytes.Buffer
	if err := nson.EncodeValue(&buf, deviceState); err != nil {
		fmt.Printf("  Encoding error: %v\n", err)
		return
	}

	data := buf.Bytes()
	fmt.Printf("  Encoded size: %d bytes\n", len(data))
	fmt.Printf("  Encoded data: %x\n", data[:min(32, len(data))])

	// 反序列化
	decoded, err := nson.DecodeValue(&buf)
	if err != nil {
		fmt.Printf("  Decoding error: %v\n", err)
		return
	}

	decodedMap, ok := decoded.(nson.Map)
	if !ok {
		fmt.Println("  Decoded value is not a map")
		return
	}

	// 验证
	if name, ok := decodedMap["name"].(nson.String); ok {
		fmt.Printf("  Decoded device name: %s\n", name)
	}

	if brightness, ok := decodedMap["brightness"].(nson.U8); ok {
		fmt.Printf("  Decoded brightness: %d/255 (%.1f%%)\n", brightness, float64(brightness)/255.0*100)
	}

	if temp, ok := decodedMap["temperature"].(nson.I16); ok {
		fmt.Printf("  Decoded temperature: %.2f°C\n", float64(temp)/100.0)
	}

	fmt.Println("  ✓ Serialization and deserialization successful!")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
