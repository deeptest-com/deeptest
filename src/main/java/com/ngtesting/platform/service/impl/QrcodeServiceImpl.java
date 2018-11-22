package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.service.intf.QrcodeService;
import org.springframework.stereotype.Service;

import java.awt.image.BufferedImage;

@Service
public class QrcodeServiceImpl extends BaseServiceImpl implements QrcodeService {

	@Override
	public JSONObject decode(String filePath) {
        BufferedImage image;
//        try {
//            image = ImageIO.read(new File(filePath));
//            LuminanceSource source = new BufferedImageLuminanceSource(image);
//            Binarizer binarizer = new HybridBinarizer(source);
//            BinaryBitmap binaryBitmap = new BinaryBitmap(binarizer);
//            Map<DecodeHintType, Object> hints = new HashMap<DecodeHintType, Object>();
//            hints.put(DecodeHintType.CHARACTER_SET, "UTF-8");
//            Result result = new MultiFormatReader().decode(binaryBitmap, hints);
//            System.out.println(result.getText());
//            JSONObject content = JSONObject.parseObject(result.getText());
//            System.out.println("encode： " + result.getBarcodeFormat());
//
//            return content;
//        } catch (Exception e) {
//            e.printStackTrace();
//        }
        return null;
	}

}
