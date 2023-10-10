import functions.RestApi;

import java.io.IOException;

public class Main {
    public static void main(String[] args) throws IOException, InterruptedException {
        RestApi data = new RestApi();
        //data.prueba();
        data.registro("prueba@gmail.com","hola1234");
    }
}