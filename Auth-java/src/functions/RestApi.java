package functions;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.io.IOException;

public class RestApi {

    final String base = "http://127.0.0.1:8050/user";

    public void registro(String email , String password) throws IOException, InterruptedException{
        String link = this.base + "/prueba";
        System.out.println(link);
        HttpClient httpClient = HttpClient.newBuilder().build();

        // Crea una solicitud GET
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(link))
                .build();

        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
        int statusCode = response.statusCode();
        System.out.println("Código de estado: " + statusCode);
        if (statusCode == 200) {
            String responseBody = response.body();
            System.out.println("Respuesta: " + responseBody);
        } else {
            System.out.println("Error al hacer la solicitud: " + response.body());
        }
    }
    }
    public void prueba() throws IOException, InterruptedException {
        String link = this.base + "/prueba";
        System.out.println(link);
        HttpClient httpClient = HttpClient.newBuilder().build();

        // Crea una solicitud GET
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(link))
                .build();

        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
        int statusCode = response.statusCode();
        System.out.println("Código de estado: " + statusCode);
        if (statusCode == 200) {
            String responseBody = response.body();
            System.out.println("Respuesta: " + responseBody);
        } else {
            System.out.println("Error al hacer la solicitud: " + response.body());
        }
    }
}
