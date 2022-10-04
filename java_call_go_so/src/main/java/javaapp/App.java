package javaapp;

public class App {
 
    public static void main(String[] args) {        
        // 启动java主Thread
        System.out.println("Java main Thread Running");        

        // 启动一个Thread运行java http server
        JavaThread jt = new JavaThread();
        jt.start();        

        // 启动一个Thread运行go http server
        GoThread gt = new GoThread();
        gt.start();        

        System.out.println("hello from java main thread");
   }

}

