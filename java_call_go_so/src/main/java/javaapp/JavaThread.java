package javaapp;

import java.io.IOException;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.eclipse.jetty.server.Server;

import org.eclipse.jetty.servlet.*;


// 启动一个线程运行java http server, 和另一个线程里的go http server互相调用.
public class JavaThread extends Thread {
   public void run() {
      try {
        System.out.println("Java Thread Running");
        JavaApp javaApp = new JavaApp();
        javaApp.Start();
      } catch(Exception e) {
         System.out.println("Java Thread Exception: "+e.getMessage());
      } finally {
         System.out.println("Java Thread Exited.");
      }
   }
}

class JavaApp {
    public void Start() throws Exception {
        int port = 8001;
        Server server = new Server(port);
        ServletHandler handler = new ServletHandler();

        // 请求全部走到这个servlet里,再根据func名作参数去调用不同函数
        handler.addServletWithMapping(FuncServlet.class, "/func");

        // 直接开启一个http proxy,代理所有http请求到go服务即可
        ServletHolder holder = new ServletHolder(new HttpProxyServlet());
        holder.setInitParameter("maxThreads", "5");
        handler.addServletWithMapping(holder, "/proxy/*");
        
        server.setHandler(handler);
        System.out.println("\nJava Http Server Started at :" + String.valueOf(port));
        server.start();
        server.join();
    }
}

@SuppressWarnings("serial")
class FuncServlet extends HttpServlet {
    @Override
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException{
        response.setStatus(HttpServletResponse.SC_OK);
        String funcName = request.getParameter("funcName");
        // call func here
        response.getWriter().println("func result="+funcName);
    }    
}

