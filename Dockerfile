# ขั้นตอนที่ 1: ใช้ Golang base image สำหรับการ build โค้ด Go
FROM golang:1.23-alpine AS build

# ตั้งค่า WORKDIR ในคอนเทนเนอร์
WORKDIR /app

# คัดลอกไฟล์ go.mod และ go.sum ก่อนเพื่อจะติดตั้ง dependencies
COPY backend/go.mod backend/go.sum ./

# ดาวน์โหลด dependencies ที่จำเป็นจาก go.mod
RUN go mod tidy

# คัดลอกไฟล์ทั้งหมดจากโปรเจ็กต์ไปยังคอนเทนเนอร์
COPY backend/ .

# สร้างแอป Go โดยใช้คำสั่ง `go build`
RUN go build -o main .

# ขั้นตอนที่ 2: ใช้ Alpine image สำหรับรันแอป Go
FROM alpine:latest

# ติดตั้ง dependencies ที่จำเป็นเพื่อรันแอป Go (เช่น ca-certificates)
RUN apk --no-cache add ca-certificates

# ตั้งค่า WORKDIR ในคอนเทนเนอร์
WORKDIR /root/

# คัดลอกไฟล์ binary ที่สร้างขึ้นจากขั้นตอนแรก
COPY --from=build /app/main .

# คัดลอกไฟล์ .env จาก root directory
COPY .env /root/.env

# ตั้งคำสั่งเริ่มต้นที่ต้องการให้รันเมื่อคอนเทนเนอร์เริ่มทำงาน
CMD ["./main"]
