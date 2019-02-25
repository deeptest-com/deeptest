package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageElementDao;
import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.service.intf.IssueFieldService;
import com.ngtesting.platform.service.intf.IssuePageElementService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Date;
import java.util.List;
import java.util.Map;

@Service
public class IssuePageElementServiceImpl extends BaseServiceImpl implements IssuePageElementService {
    Log logger = LogFactory.getLog(IssueJqlColumnServiceImpl.class);

	@Autowired
    IssuePageElementDao elementDao;

    @Autowired
    IssueFieldService fieldService;

    @Override
    @Transactional
    public void saveAll(Integer orgId, Integer pageId, List<Map> maps) {
        elementDao.removeOthers(maps, pageId, orgId);

        if (maps.size() > 0) {
            int ordr = 1;
            for (Map map: maps) {
                map.put("ordr", ordr++);

                Object id = map.get("id");
                String key = map.get("key").toString();

                if (id == null) {
                    IsuField field = fieldService.getField(key, orgId);

                    IsuPageElement elm = new IsuPageElement(field.getColCode(), field.getLabel(),
                            field.getType(), field.getInput(),
                            field.getFullLine(), field.getRequired(), field.getReadonly(), field.getBuildIn(),
                            field.getKey(), field.getFieldId(), pageId, orgId, ordr);

                    if (elm.getInput().equals("textarea") || elm.getInput().equals("richtext")) {
                        elm.setFullLine(Boolean.TRUE);
                    }
                    elementDao.save(elm);
                    map.put("id", elm.getId().toString());
                }
            }

            long start = new Date().getTime();
            elementDao.saveOrdrs(maps, pageId, orgId);
            long end = new Date().getTime();

            logger.info("Update ordrs for " + maps.size() + " records spend " + (end - start) + " milliseconds");
        }
    }

    @Override
    public void updateProp(Integer id, String prop, String val, Integer orgId) {
        if ("required".equals(prop) || "fullLine".equals(prop) || "readonly".equals(prop)) {

        } else {
            val = "'" + val + "'";
        }
        elementDao.updateProp(id, prop, val, orgId);
    }

//    @Override
//    public void save(IsuPageElement element) {
//        Integer maxOrder = elementDao.getMaxFieldOrdr(element.getTabId());
//        maxOrder = maxOrder == null? 0: maxOrder;
//        element.setOrdr(maxOrder + 1);
//        elementDao.save(element);
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
