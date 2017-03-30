package com.ngtesting.platform.service.impl;

import java.awt.image.BufferedImage;
import java.io.File;
import java.util.HashMap;
import java.util.Map;

import javax.imageio.ImageIO;

import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;
import com.google.zxing.Binarizer;
import com.google.zxing.BinaryBitmap;
import com.google.zxing.DecodeHintType;
import com.google.zxing.LuminanceSource;
import com.google.zxing.MultiFormatReader;
import com.google.zxing.Result;
import com.google.zxing.client.j2se.BufferedImageLuminanceSource;
import com.google.zxing.common.HybridBinarizer;
import com.ngtesting.platform.service.QrcodeService;

@Service
public class QrcodeServiceImpl extends BaseServiceImpl implements QrcodeService {

	@Override
	public JSONObject decode(String filePath) { 
        BufferedImage image;  
        try {  
            image = ImageIO.read(new File(filePath));  
            LuminanceSource source = new BufferedImageLuminanceSource(image);  
            Binarizer binarizer = new HybridBinarizer(source);  
            BinaryBitmap binaryBitmap = new BinaryBitmap(binarizer);  
            Map<DecodeHintType, Object> hints = new HashMap<DecodeHintType, Object>();  
            hints.put(DecodeHintType.CHARACTER_SET, "UTF-8");  
            Result result = new MultiFormatReader().decode(binaryBitmap, hints);
            System.out.println(result.getText());
            JSONObject content = JSONObject.parseObject(result.getText());
            System.out.println("encode： " + result.getBarcodeFormat()); 
            
            return content;
        } catch (Exception e) {  
            e.printStackTrace();  
        } 
        return null;
	}
	
}
