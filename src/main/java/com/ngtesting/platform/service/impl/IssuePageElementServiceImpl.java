package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.IssuePageElementDao;
import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.service.IssueFieldService;
import com.ngtesting.platform.service.IssuePageElementService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class IssuePageElementServiceImpl extends BaseServiceImpl implements IssuePageElementService {

	@Autowired
    IssuePageElementDao elementDao;

    @Autowired
    IssueFieldService fieldService;

    @Override
    @Transactional
    public void saveAll(Integer orgId, Integer pageId, Integer tabId, JSONArray jsonArr) {
        for (Object obj: jsonArr) {
            JSONObject json = (JSONObject) obj;

            Integer id = json.getInteger("id");
            String key = json.getString("key");

            if (id == null) {
                IsuField field = fieldService.getField(key);

                IsuPageElement elm = new IsuPageElement(field.getCode(), field.getLabel(),
                        field.getType(), field.getInput(), field.getFullLine(), field.getRequired(),
                        field.getKey(), field.getFieldId(), tabId, pageId, orgId);

                elementDao.save(elm);
            }
        }
    }

    @Override
    public void updateProp(String id, String prop, String val, Integer orgId) {
        if ("required".equals(prop) || "fullLine".equals(prop)) {

        } else {
            val = "'" + val + "'";
        }
        elementDao.updateProp(id, prop, val, orgId);
    }

//    @Override
//    public void add(IsuPageElement element) {
//        Integer maxOrder = elementDao.getMaxFieldOrdr(element.getTabId());
//        maxOrder = maxOrder == null? 0: maxOrder;
//        element.setOrdr(maxOrder + 1);
//        elementDao.add(element);
//    }
//
//    @Override
//    public boolean remove(Integer id, Integer orgId) {
//        Integer count = elementDao.remove(id, orgId);
//        if (count == 0) {
//            return false;
//        }
//
//        return true;
//    }

}
