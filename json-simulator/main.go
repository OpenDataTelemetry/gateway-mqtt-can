package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"

	// "math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	PAD "github.com/pinpt/go-common/strings"
)

func hexToB64(h string) (string, error) {
	p, err := hex.DecodeString(h)
	if err != nil {
		log.Fatal(err)
	}
	b64 := base64.StdEncoding.EncodeToString([]byte(p))
	return b64, err
}

func randSensorData(sensorType string, i int) string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder

	switch sensorType {
	case "temperature":
		sb.WriteString("01")
		min := 0
		max := 400
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "humidity":
		sb.WriteString("02")
		min := 0
		max := 100
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "boardVoltage":
		sb.WriteString("0c")
		min := 0
		max := 4200
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "counter":
		sb.WriteString("0b")
		min := 0
		// max := 16777215
		v := PAD.PadLeft((strconv.FormatUint(uint64(min+i), 16)), 6, 0x30)
		sb.WriteString(v)

	case "counter_0d":
		sb.WriteString("0d")
		min := 0
		max := 4096
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "distance":
		sb.WriteString("13")
		min := 0
		max := 4096
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)
		// case 200:
		// 	jsonProtocolParser()

		// case 3:
		// 	khompProtocolParser()

		// case 4:
		// 	khompProtocolParser()
	}
	return sb.String()
	// return sb.String()
}

func randDeviceData(deviceType string, i int) (string, string) {
	type deviceInfo struct {
		nodeName string
		devEUI   string
	}
	var sb strings.Builder
	var stb strings.Builder

	// sb.WriteString(`{"applicationID":"`)

	switch deviceType {
	case "PD01_DCU":
		// var sdb strings.Builder
		// var stb strings.Builder
		// deviceInfos := []deviceInfo{
		// 	{nodeName: "ECU_1", devEUI: "0001"},
		// }
		// j := float64(i)
		// if i >= len(deviceInfos) {
		// 	j = math.Mod(float64(i), float64(len(deviceInfos)))
		// }
		// k := uint(j)

		// sb.WriteString(`6","applicationName":"smartcampusmaua","nodeName":"`)
		// sb.WriteString(deviceInfos[k].nodeName)
		// sb.WriteString(`","devEUI":"`)
		// sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`{ "fields":{"sdl3_battery_voltage":12.4,"engine_rpm":7000,"gps_speed":84,"gear":4},"name":"pd01_dcu","timestamp":1713288706239}`)
		// sdb.WriteString(randSensorData("counter", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("temperature", i))
		// sdb.WriteString(randSensorData("humidity", i))
		// sdb.WriteString(randSensorData("boardVoltage", i))
		// data, _ := hexToB64(sdb.String())
		// sb.WriteString(data)
		// sb.WriteString(`"}`)
		stb.WriteString(`FSAELive/IC/MauaRacing/Telemetry/rx`)
		// stb.WriteString(deviceInfos[k].devEUI)
		// stb.WriteString(`/rx`)

	case "PD02_DCU":
		// var sdb strings.Builder
		// var stb strings.Builder
		// deviceInfos := []deviceInfo{
		// 	{nodeName: "ECU_1", devEUI: "0001"},
		// }
		// j := float64(i)
		// if i >= len(deviceInfos) {
		// 	j = math.Mod(float64(i), float64(len(deviceInfos)))
		// }
		// k := uint(j)

		// sb.WriteString(`6","applicationName":"smartcampusmaua","nodeName":"`)
		// sb.WriteString(deviceInfos[k].nodeName)
		// sb.WriteString(`","devEUI":"`)
		// sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`{"fields":{"engine_oil_pressure":3.0,"engine_temperature":70,"engine_intake_air_temperature":84,"engine_ecu_temperature":4},"name":"pd02_dcu","timestamp":1713288706239}`)
		// sdb.WriteString(randSensorData("counter", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("temperature", i))
		// sdb.WriteString(randSensorData("humidity", i))
		// sdb.WriteString(randSensorData("boardVoltage", i))
		// data, _ := hexToB64(sdb.String())
		// sb.WriteString(data)
		// sb.WriteString(`"}`)
		stb.WriteString(`FSAELive/IC/MauaRacing/Telemetry/rx`)
		// stb.WriteString(deviceInfos[k].devEUI)
		// stb.WriteString(`/rx`)

	case "PD03_DCU":
		// var sdb strings.Builder
		// var stb strings.Builder
		// deviceInfos := []deviceInfo{
		// 	{nodeName: "ECU_1", devEUI: "0001"},
		// }
		// j := float64(i)
		// if i >= len(deviceInfos) {
		// 	j = math.Mod(float64(i), float64(len(deviceInfos)))
		// }
		// k := uint(j)

		// sb.WriteString(`6","applicationName":"smartcampusmaua","nodeName":"`)
		// sb.WriteString(deviceInfos[k].nodeName)
		// sb.WriteString(`","devEUI":"`)
		// sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`{"fields":{"brake_pressure_front":120,"brake_pressure_rear":250,"manifold_air_pressure":1,"fuel_pressure":3.5},"name":"pd03_dcu","timestamp":1713288706239}`)
		// sdb.WriteString(randSensorData("counter", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("temperature", i))
		// sdb.WriteString(randSensorData("humidity", i))
		// sdb.WriteString(randSensorData("boardVoltage", i))
		// data, _ := hexToB64(sdb.String())
		// sb.WriteString(data)
		// sb.WriteString(`"}`)
		stb.WriteString(`FSAELive/IC/MauaRacing/Telemetry/rx`)
		// stb.WriteString(deviceInfos[k].devEUI)
		// stb.WriteString(`/rx`)

	case "PD04_DCU":
		// var sdb strings.Builder
		// var stb strings.Builder
		// deviceInfos := []deviceInfo{
		// 	{nodeName: "ECU_1", devEUI: "0001"},
		// }
		// j := float64(i)
		// if i >= len(deviceInfos) {
		// 	j = math.Mod(float64(i), float64(len(deviceInfos)))
		// }
		// k := uint(j)

		// sb.WriteString(`6","applicationName":"smartcampusmaua","nodeName":"`)
		// sb.WriteString(deviceInfos[k].nodeName)
		// sb.WriteString(`","devEUI":"`)
		// sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`{"fields":{"exhaust_temperature_cylinder_1":720,"exhaust_temperature_cylinder_1":750,"exhaust_temperature_cylinder_1":700,"sdl3_temperature":27.5},"name":"pd04_dcu","timestamp":1713288706239}`)
		// sdb.WriteString(randSensorData("counter", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("temperature", i))
		// sdb.WriteString(randSensorData("humidity", i))
		// sdb.WriteString(randSensorData("boardVoltage", i))
		// data, _ := hexToB64(sdb.String())
		// sb.WriteString(data)
		// sb.WriteString(`"}`)
		stb.WriteString(`FSAELive/IC/MauaRacing/Telemetry/rx`)
		// stb.WriteString(deviceInfos[k].devEUI)
		// stb.WriteString(`/rx`)

	case "PD05_DCU":
		// var sdb strings.Builder
		// var stb strings.Builder
		// deviceInfos := []deviceInfo{
		// 	{nodeName: "ECU_1", devEUI: "0001"},
		// }
		// j := float64(i)
		// if i >= len(deviceInfos) {
		// 	j = math.Mod(float64(i), float64(len(deviceInfos)))
		// }
		// k := uint(j)

		// sb.WriteString(`6","applicationName":"smartcampusmaua","nodeName":"`)
		// sb.WriteString(deviceInfos[k].nodeName)
		// sb.WriteString(`","devEUI":"`)
		// sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`{"fields":{"lambda_1":0.8,"acc_x":2.5,"axx_y":2.5,"acc_z":9.8},"name":"pd05_dcu","timestamp":1713288706239}`)
		// sdb.WriteString(randSensorData("counter", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("temperature", i))
		// sdb.WriteString(randSensorData("humidity", i))
		// sdb.WriteString(randSensorData("boardVoltage", i))
		// data, _ := hexToB64(sdb.String())
		// sb.WriteString(data)
		// sb.WriteString(`"}`)
		stb.WriteString(`FSAELive/IC/MauaRacing/Telemetry/rx`)
		// stb.WriteString(deviceInfos[k].devEUI)
		// stb.WriteString(`/rx`)

	case "PD06_DCU":
		// var sdb strings.Builder
		// var stb strings.Builder
		// deviceInfos := []deviceInfo{
		// 	{nodeName: "ECU_1", devEUI: "0001"},
		// }
		// j := float64(i)
		// if i >= len(deviceInfos) {
		// 	j = math.Mod(float64(i), float64(len(deviceInfos)))
		// }
		// k := uint(j)

		// sb.WriteString(`6","applicationName":"smartcampusmaua","nodeName":"`)
		// sb.WriteString(deviceInfos[k].nodeName)
		// sb.WriteString(`","devEUI":"`)
		// sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`{"fields":{"throttle_position":0.8,"fuel_used":2.5,"beacon":2,"gps_altitude":770},"name":"pd06_dcu","timestamp":1713288706239}`)
		// sdb.WriteString(randSensorData("counter", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("temperature", i))
		// sdb.WriteString(randSensorData("humidity", i))
		// sdb.WriteString(randSensorData("boardVoltage", i))
		// data, _ := hexToB64(sdb.String())
		// sb.WriteString(data)
		// sb.WriteString(`"}`)
		stb.WriteString(`FSAELive/IC/MauaRacing/Telemetry/rx`)
		// stb.WriteString(deviceInfos[k].devEUI)
		// stb.WriteString(`/rx`)

	case "PD07_DCU":
		// var sdb strings.Builder
		// var stb strings.Builder
		// deviceInfos := []deviceInfo{
		// 	{nodeName: "ECU_1", devEUI: "0001"},
		// }
		// j := float64(i)
		// if i >= len(deviceInfos) {
		// 	j = math.Mod(float64(i), float64(len(deviceInfos)))
		// }
		// k := uint(j)

		// sb.WriteString(`6","applicationName":"smartcampusmaua","nodeName":"`)
		// sb.WriteString(deviceInfos[k].nodeName)
		// sb.WriteString(`","devEUI":"`)
		// sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`{"fields":{"gps_latitude":-23.844756,"gps_longitude":-46.23548},"name":"pd07_dcu","timestamp":1713288706239}`)
		// sdb.WriteString(randSensorData("counter", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("counter_0d", i))
		// sdb.WriteString(randSensorData("temperature", i))
		// sdb.WriteString(randSensorData("humidity", i))
		// sdb.WriteString(randSensorData("boardVoltage", i))
		// data, _ := hexToB64(sdb.String())
		// sb.WriteString(data)
		// sb.WriteString(`"}`)
		stb.WriteString(`FSAELive/IC/MauaRacing/Telemetry/rx`)
		// stb.WriteString(deviceInfos[k].devEUI)
		// stb.WriteString(`/rx`)

	}
	return stb.String(), sb.String()

}

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(1000 * time.Second)
		done <- true
	}()

	id := uuid.New().String()
	var sbMqtt strings.Builder
	sbMqtt.WriteString("mqtt-")
	sbMqtt.WriteString(id)
	i := 0

	// MQTT
	// topic := "application/6/node/0004a30b00e94314/rx"
	broker := "mqtt://mqtt.maua.br:1883"
	mqttClientId := sbMqtt.String()
	qos := 0
	password := "public"
	user := "public"

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(mqttClientId)
	opts.SetUsername(user)
	opts.SetPassword(password)

	choke := make(chan [2]string)

	opts.SetDefaultPublishHandler(func(mqttClient MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	mqttClient := MQTT.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			mqttClient.Disconnect(250)
			fmt.Println("Sample Publisher Disconnected")
			return
		case t := <-ticker.C:
			// fmt.Printf("\n%s,\nSent payload: %v\n\n", t, payload[rand.Intn(len(payload))])
			topic, payload := randDeviceData("PD01_DCU", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("PD02_DCU", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("PD03_DCU", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("PD04_DCU", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("PD05_DCU", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("PD06_DCU", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("PD07_DCU", i)
			mqttClient.Publish(topic, byte(qos), false, payload)

			fmt.Printf("\nt: %v\n\n", t.UnixNano())
			i++
			// token.Wait()

		}
	}
}
