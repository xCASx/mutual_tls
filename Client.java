import javax.net.ssl.HttpsURLConnection;
import javax.net.ssl.SSLSocketFactory;
import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.URL;

/**
 * @author Maxim Kolesnikov <swe.kolesnikov@gmail.com>
 */
public class Client {
    public static void main(String[] args) throws Exception {
        SSLSocketFactory sslsocketfactory = (SSLSocketFactory) SSLSocketFactory.getDefault();
        URL url = new URL("https://localhost:8443");
        HttpsURLConnection conn = (HttpsURLConnection)url.openConnection();
        conn.setSSLSocketFactory(sslsocketfactory);
        InputStream inputstream = conn.getInputStream();
        InputStreamReader inputstreamreader = new InputStreamReader(inputstream);
        BufferedReader bufferedreader = new BufferedReader(inputstreamreader);

        String string;
        while ((string = bufferedreader.readLine()) != null) {
            System.out.println("Received " + string);
        }
    }
}
