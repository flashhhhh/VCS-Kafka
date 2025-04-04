# Kafka

##  Khái niệm
Apache Kafka là một nền tảng phân tán xử lý luồng dữ liệu cho việc nhập dữ liệu và xử lý dữ liệu thời gian thực. Streaming data là dữ liệu được sinh ra liên tục từ hàng nghìn nguồn dữ liệu, và dữ liệu thường được gửi đi ngay lập tức. Một nền tảng streamming cần xử lý với khối lượng dữ liệu khổng lồ này, và xử lý chúng tuần tự và liên tục.

Kafka cung cấp 3 chức năng chính cho người dùng:
+ Publish và subcribe tới luồng dữ liệu
+ Lưu trữ luồng dữ liệu hiệu quả theo thứ tự mà dữ liệu được sinh ra
+ Xử lý các luồng dữ liệu trong thời gian thực

## Mục đích sử dụng
### Gửi và nhận tin nhắn
Kafka có thể thay thế hiệu quả cho các message broker truyền thống khác. Message brokers được sử dụng trong nhiều trường hợp (tách biệt việc xử lý dữ liệu khỏi nguồn dữ liệu, làm bộ đệm lưu trữ các tin nhắn chưa được xử lý, ...).

So với hầu hết các hệ thống messaging, Kafka có thông lượng tốt hơn, có các cơ chế như các partition, replication giúp khả năng xử lý với mất mát dữ liệu tốt hơn, giúp Kafka trở thành giải pháp hiệu quả cho các ứng dụng yêu cầu xử lý thông lượng tin nhắn lớn.

### Metrics
Kafka có thể được dùng để giám sát hoạt động của ứng dụng từ dữ liệu của ứng dụng đó. Điều này bao gồm tổng hợp các thống kê từ các ứng dụng bị phân tán ở nhiều chỗ về một nguồn cấp dữ liệu hoạt động tập trung.

### Tổng hợp log
Tổng hợp log thường thu thập các log files từ servers và lưu trữ ở một nơi tập trung để xử lý. Kafka tóm tắt các chi tiết của các tệp và cung cấp một bản tóm tắt gọn hơn về dữ liệu nhật ký hoặc sự kiện dưới dạng một luồng tin nhắn, cho phép xử lý dữ liệu ít độ trễ hơn và hỗ trợ tốt hơn cho nhiều nguồn dữ liệu và nhiều nguồn tiêu thụ phi tập trung.

So với các hệ thống ghi log tập trung như Scribe hoặc Flume, Kafka cung cấp hiệu suất tốt như nhau, đảm bảo dữ liệu không bị mất do khả năng replication và độ trễ đầu cuối thấp hơn nhiều.

### Stream Processing
So với việc xử lý dữ liệu theo batch yêu cầu dữ liệu đầu vào phải là một khối, khiến cho việc xử lý dữ liệu không được liên tục, Kafka cho phép dữ liệu mới được sinh ra được truyền trực tiếp đến nơi xử lý dữ liệu, giúp việc nhập hay xử lý dữ liệu được diễn ra trong thời gian thực

### Kiến trúc Event-Driven
Dùng Kafka làm trung gian để các dịch vụ có thể giao tiếp mà không phụ thuộc trực tiếp vào nhau.

## So sánh điểm mạnh, điểm yếu với các Message Queue khác
# Kafka

##  Khái niệm
Apache Kafka là một nền tảng phân tán xử lý luồng dữ liệu cho việc nhập dữ liệu và xử lý dữ liệu thời gian thực. Streaming data là dữ liệu được sinh ra liên tục từ hàng nghìn nguồn dữ liệu, và dữ liệu thường được gửi đi ngay lập tức. Một nền tảng streamming cần xử lý với khối lượng dữ liệu khổng lồ này, và xử lý chúng tuần tự và liên tục.

Kafka cung cấp 3 chức năng chính cho người dùng:
+ Publish và subcribe tới luồng dữ liệu
+ Lưu trữ luồng dữ liệu hiệu quả theo thứ tự mà dữ liệu được sinh ra
+ Xử lý các luồng dữ liệu trong thời gian thực

## Mục đích sử dụng
### Gửi và nhận tin nhắn
Kafka có thể thay thế hiệu quả cho các message broker truyền thống khác. Message brokers được sử dụng trong nhiều trường hợp (tách biệt việc xử lý dữ liệu khỏi nguồn dữ liệu, làm bộ đệm lưu trữ các tin nhắn chưa được xử lý, ...).

So với hầu hết các hệ thống messaging, Kafka có thông lượng tốt hơn, có các cơ chế như các partition, replication giúp khả năng xử lý với mất mát dữ liệu tốt hơn, giúp Kafka trở thành giải pháp hiệu quả cho các ứng dụng yêu cầu xử lý thông lượng tin nhắn lớn.

### Metrics
Kafka có thể được dùng để giám sát hoạt động của ứng dụng từ dữ liệu của ứng dụng đó. Điều này bao gồm tổng hợp các thống kê từ các ứng dụng bị phân tán ở nhiều chỗ về một nguồn cấp dữ liệu hoạt động tập trung.

### Tổng hợp log
Tổng hợp log thường thu thập các log files từ servers và lưu trữ ở một nơi tập trung để xử lý. Kafka tóm tắt các chi tiết của các tệp và cung cấp một bản tóm tắt gọn hơn về dữ liệu nhật ký hoặc sự kiện dưới dạng một luồng tin nhắn, cho phép xử lý dữ liệu ít độ trễ hơn và hỗ trợ tốt hơn cho nhiều nguồn dữ liệu và nhiều nguồn tiêu thụ phi tập trung.

So với các hệ thống ghi log tập trung như Scribe hoặc Flume, Kafka cung cấp hiệu suất tốt như nhau, đảm bảo dữ liệu không bị mất do khả năng replication và độ trễ đầu cuối thấp hơn nhiều.

### Stream Processing
So với việc xử lý dữ liệu theo batch yêu cầu dữ liệu đầu vào phải là một khối, khiến cho việc xử lý dữ liệu không được liên tục, Kafka cho phép dữ liệu mới được sinh ra được truyền trực tiếp đến nơi xử lý dữ liệu, giúp việc nhập hay xử lý dữ liệu được diễn ra trong thời gian thực

### Kiến trúc Event-Driven
Dùng Kafka làm trung gian để các dịch vụ có thể giao tiếp mà không phụ thuộc trực tiếp vào nhau.

# Kafka

##  Khái niệm
Apache Kafka là một nền tảng phân tán xử lý luồng dữ liệu cho việc nhập dữ liệu và xử lý dữ liệu thời gian thực. Streaming data là dữ liệu được sinh ra liên tục từ hàng nghìn nguồn dữ liệu, và dữ liệu thường được gửi đi ngay lập tức. Một nền tảng streamming cần xử lý với khối lượng dữ liệu khổng lồ này, và xử lý chúng tuần tự và liên tục.

Kafka cung cấp 3 chức năng chính cho người dùng:
+ Publish và subcribe tới luồng dữ liệu
+ Lưu trữ luồng dữ liệu hiệu quả theo thứ tự mà dữ liệu được sinh ra
+ Xử lý các luồng dữ liệu trong thời gian thực

## Mục đích sử dụng
### Gửi và nhận tin nhắn
Kafka có thể thay thế hiệu quả cho các message broker truyền thống khác. Message brokers được sử dụng trong nhiều trường hợp (tách biệt việc xử lý dữ liệu khỏi nguồn dữ liệu, làm bộ đệm lưu trữ các tin nhắn chưa được xử lý, ...).

So với hầu hết các hệ thống messaging, Kafka có thông lượng tốt hơn, có các cơ chế như các partition, replication giúp khả năng xử lý với mất mát dữ liệu tốt hơn, giúp Kafka trở thành giải pháp hiệu quả cho các ứng dụng yêu cầu xử lý thông lượng tin nhắn lớn.

### Metrics
Kafka có thể được dùng để giám sát hoạt động của ứng dụng từ dữ liệu của ứng dụng đó. Điều này bao gồm tổng hợp các thống kê từ các ứng dụng bị phân tán ở nhiều chỗ về một nguồn cấp dữ liệu hoạt động tập trung.

### Tổng hợp log
Tổng hợp log thường thu thập các log files từ servers và lưu trữ ở một nơi tập trung để xử lý. Kafka tóm tắt các chi tiết của các tệp và cung cấp một bản tóm tắt gọn hơn về dữ liệu nhật ký hoặc sự kiện dưới dạng một luồng tin nhắn, cho phép xử lý dữ liệu ít độ trễ hơn và hỗ trợ tốt hơn cho nhiều nguồn dữ liệu và nhiều nguồn tiêu thụ phi tập trung.

So với các hệ thống ghi log tập trung như Scribe hoặc Flume, Kafka cung cấp hiệu suất tốt như nhau, đảm bảo dữ liệu không bị mất do khả năng replication và độ trễ đầu cuối thấp hơn nhiều.

### Stream Processing
So với việc xử lý dữ liệu theo batch yêu cầu dữ liệu đầu vào phải là một khối, khiến cho việc xử lý dữ liệu không được liên tục, Kafka cho phép dữ liệu mới được sinh ra được truyền trực tiếp đến nơi xử lý dữ liệu, giúp việc nhập hay xử lý dữ liệu được diễn ra trong thời gian thực

### Kiến trúc Event-Driven
Dùng Kafka làm trung gian để các dịch vụ có thể giao tiếp mà không phụ thuộc trực tiếp vào nhau.

| Tính chất    | Apache Kafka | RabbitMQ |
| -------- | ------- | ------ |
| Kiến trúc | **Kafka sử dụng mô hình nhật ký phân vùng, nên có thể kết hợp theo cả 2 hướng message queue và publish subcribe** | RabbitMQ sử dụng message queue |
| Khả năng mở rộng | **Cho phép các partition được phân tán trên nhiều server khác nhau** | Tăng số lượng consumer để giảm bớt công việc cho các consumer khác |
| Thời gian lưu trữ tin nhắn | **Theo quy tắc được định nghĩa, ví dụ như 1 ngày, hoặc theo config của người dùng** | Theo ack, tin nhắn sẽ bị xóa ngay lập tức khi có 1 consumer tiêu thụ thành công |
| Nhiều consumer | **Nhiều consumer có thể subcribe tới cùng 1 topic, do Kafka cho phép cùng 1 tin nhắn có thể được gửi lại cho các consumer khác trong 1 khoảng thời gian** | Nhiều consumer không thể tiêu thụ cùng 1 tin nhắn, do tin nhắn được xóa ngay lập tức khi nó được tiêu thụ |
| Replication | Các topic được mặc định là nhân bản, nhưng người dùng có thể setup để không nhân bản cho topic đó | Tin nhắn được mặc định là không nhân bản, nhưng người dùng có thể setup để nhân bản |
| Thứ tự tin nhắn | **Các consumer nhận được tin nhắn theo đúng thứ tự đầu vào do kiến trúc nhật ký phân vùng** | Tin nhắn được chuyển đến cho consumer theo thứ tự đến hàng đợi. Nếu có nhiều consumer, mỗi consumer sẽ xử lý một tập con của tin nhắn trong hàng đợi |
| Giao thức | Kafka sử dụng giao thức nhị phân trên TCP | **Giao thức hàng đợi nhắn tin nâng cao (AMQP) với hỗ trợ thông qua các plugin: MQTT, Stomp**|