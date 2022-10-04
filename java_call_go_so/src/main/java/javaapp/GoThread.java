package javaapp;

import com.sun.jna.*;


// 启动一个java Thread，运行go编译的.so库里的httpserver
public class GoThread extends Thread {
   public void run() {      
      try {
         System.out.println("Go Thread Running");
         GoApp goApp = new GoApp();
         goApp.Start();
      } catch(Exception e) {
         System.out.println("Go Thread Exception: "+e.getMessage());
      } finally {
         System.out.println("Go Thread Exited.");
      }
   }
}

// go app 入口
class GoApp {

   public interface LibGo extends Library {
        Pointer Hello(Pointer a);  //传参方式调用
        void HttpServer();  //启动httpserver方式走http调用
   }

   public void Start() {
      // java使用jna加载so库,调用go函数
      String libgoPath = "/home/michael/github/java/sample-code/call_go_so/go/libgo.so";
      LibGo libgo = Native.loadLibrary(libgoPath, LibGo.class);

      //在jvm之外的jna内存中开辟大小为size的byte[],这是专门给jna调用so库使用的内存
      //注意size大小要比真实字符串长度大1,否则在赋值的时候会提示数组边界越界
      //jna会申请一块内存用于和so库交换数据,大小为5byte,从位置0开始写入字符
      //TODO: 申请的内存不会被自动回收,会内存泄漏,需要手动释放: Native.free(Pointer.nativeValue(p))
      System.out.println("GoThread Call go func: Hello");
      Pointer p1 = new Memory(5);
      p1.setString(0, "abc"); 
      Pointer p2 = libgo.Hello(p1); //调用Hello函数,传参字符串
      String s = p2.getString(0);
      System.out.println("go returns = "+s);  

      System.out.println("GoThread Call go func: HttpServer");
      libgo.HttpServer(); //调用HttpServer函数,启动一个go http server
   }
}

