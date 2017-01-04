package cn.mobiu.events.util;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class FileUtils {
    public static String ReadFile(String fullFileName) {
        File file = new File(fullFileName);
        Scanner scanner = null;
        StringBuilder buffer = new StringBuilder();
        try {
            scanner = new Scanner(file, "utf-8");
            while (scanner.hasNextLine()) {
                buffer.append(scanner.nextLine());
            }
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } finally {
            if (scanner != null) {
                scanner.close();
            }
        }

        return buffer.toString();
    }
    
    public static void CreateDirIfNeeded(String dir) {
	  File file = new File(dir);
	  if(!file.exists()){
		  file.mkdirs();
	  }
    }
    
}