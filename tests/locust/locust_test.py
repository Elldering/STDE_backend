from locust import HttpUser, TaskSet, task, between, events
from faker import Faker
import random
import json
import matplotlib.pyplot as plt
from PIL import Image
import statistics

fake = Faker()

class ApiUserTasks(TaskSet):
    wait_time = between(1, 3)

    @task(1)
    def get_buckets(self):
        self.execute_test("GET Buckets", "/api/test-routes/s3/buckets", "GET")

    @task(2)
    def upload_file(self):
        files = {"file": ("test_file.txt", "Hello, World!", "text/plain")}
        key = f"/uploads/{fake.uuid4()}.txt"
        self.execute_test("Upload File", f"/api/test-routes/s3/upload{key}", "POST", files=files)

    @task(1)
    def delete_file(self):
        key = f"/uploads/{fake.uuid4()}.txt"
        self.execute_test("Delete File", f"/api/test-routes/s3/files{key}", "DELETE")

    @task(3)
    def get_reviews(self):
        self.execute_test("GET Reviews", "/api/test-routes/reviews/", "GET")

    @task(2)
    def post_review(self):
        data = {
            "auth_user_sender_id": 108,
            "auth_user_recipient_id": 108,
            "grade": random.randint(1, 5),
            "comment": fake.text(max_nb_chars=150),
        }
        self.execute_test("POST Review", "/api/test-routes/reviews/", "POST", data=json.dumps(data))

    @task(1)
    def get_user_group(self):
        group_id = random.randint(1, 20)  
        self.execute_test("GET User Group", f"/api/test-routes/user-groups/{group_id}", "GET")

    @task(2)
    def create_user_group(self):
        data = {"name": fake.company()}
        self.execute_test("POST User Group", "/api/test-routes/user-groups/", "POST", data=json.dumps(data))

    def execute_test(self, test_name, url, method, **kwargs):
        response = None
        try:
            if method == "GET":
                response = self.client.get(url)
            elif method == "POST":
                response = self.client.post(url, **kwargs, headers={"Content-Type": "application/json"})
            elif method == "DELETE":
                response = self.client.delete(url)

            status = "✔️ Успех" if response.status_code == 200 else f"❌ Ошибка ({response.status_code})"
            print(f"[{test_name}] → {status}")

            if response.status_code == 200:
                response_times.append(response.elapsed.total_seconds() * 1000)  # В миллисекундах
                successful_requests.append(test_name)
            else:
                failed_requests.append(test_name)
                failed_response_times.append(response.elapsed.total_seconds() * 1000)

        except Exception as e:
            print(f"[{test_name}] ❌ Ошибка: {e}")
            failed_requests.append(test_name)




class ApiUser(HttpUser):
    host = "http://localhost:8085"
    tasks = [ApiUserTasks]

response_times = []
failed_response_times = []
successful_requests = []
failed_requests = []

# @events.quitting.add_listener
# def on_quitting(environment, **kwargs):
#     print("\n========== Итоговый отчет по нагрузочному тестированию ==========")
#     print(f"Успешных запросов: {len(successful_requests)}")
#     print(f"Ошибок: {len(failed_requests)}")
    
#     if response_times:
#         avg_time = statistics.mean(response_times)
#         print(f"Среднее время ответа: {avg_time:.2f} ms")
#         print(f"Максимальное время ответа: {max(response_times):.2f} ms")
#         print(f"Минимальное время ответа: {min(response_times):.2f} ms")

#     plt.figure(figsize=(14, 8))

#     plt.subplot(2, 1, 1)
#     plt.plot(response_times, label="Успешные запросы", color="blue")
#     plt.plot(failed_response_times, label="Неуспешные запросы", color="red")
#     plt.xlabel("Номер запроса")
#     plt.ylabel("Время ответа (мс)")
#     plt.title("Время отклика запросов")
#     plt.legend()
#     plt.grid(True)

#     plt.subplot(2, 1, 2)
#     labels = ["Успешные", "Неуспешные"]
#     values = [len(successful_requests), len(failed_requests)]
#     plt.bar(labels, values, color=["green", "red"])
#     plt.title("Успехи / Ошибки")
#     plt.ylabel("Количество запросов")

#     plt.tight_layout()
#     plt.savefig("test_results.png")
#     plt.close()

#     try:
#         img = Image.open("test_results.png")
#         img.show()
#     except:
#         print("📊 График сохранен как test_results.png")
