package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueFieldDao;
import com.ngtesting.platform.model.IsuFieldDefine;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import com.ngtesting.platform.service.intf.IssueJqlColumnService;
import com.ngtesting.platform.service.intf.UserService;
import com.ngtesting.platform.utils.StringUtil;
import com.ngtesting.platform.vo.IsuJqlColumn;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StringUtils;

import java.util.*;

@Service
public class IssueJqlColumnServiceImpl extends BaseServiceImpl implements IssueJqlColumnService {
    Log logger = LogFactory.getLog(IssueJqlColumnServiceImpl.class);

    @Autowired
    UserService userService;

    @Autowired
    IssueFieldDao fieldDao;

    @Autowired
    IssueDynamicFormService dynamicFormService;

    @Override
    @Transactional
    public List<IsuJqlColumn> loadColumns(TstUser user) {
        List<IsuJqlColumn> ret = new LinkedList<>();

        String columnsStr = user.getIssueColumns();
        if (StringUtils.isEmpty(columnsStr) || columnsStr.indexOf("null") > -1) {
            columnsStr = buildDefaultColStr(user);
        }
        List<String> colShowArr = new ArrayList<>(Arrays.asList(columnsStr.split(",")));

        List<Map> fields = dynamicFormService.fetchOrgField(
                user.getDefaultOrgId(), user.getDefaultPrjId(), "col");

        Map<String, Map> colAllMap = new HashMap<>();

        for (Map field : fields) {
            String code = field.get("colCode").toString();
            colAllMap.put(code, field);
        }

        for (String code : colShowArr) {
            if (!colAllMap.containsKey(code)) { // 删除了
                continue;
            }

            colAllMap.get(code).put("defaultShowInColumns", true);
            IsuJqlColumn vo = new IsuJqlColumn(colAllMap.get(code));
            ret.add(vo);

            colAllMap.put(code, null);
        }

        for (String code : colAllMap.keySet()) {
            if (colAllMap.get(code) == null) { // 已添加了
                continue;
            }

            colAllMap.get(code).put("defaultShowInColumns", false);
            IsuJqlColumn vo = new IsuJqlColumn(colAllMap.get(code));
            ret.add(vo);
        }

        return ret;
    }

    @Override
    @Transactional
    public String buildDefaultColStr(TstUser user) {
        String ret = "";

        List<IsuFieldDefine> cols = fieldDao.listDefaultFieldInColumns();
        int i = 0;
        for (IsuFieldDefine col : cols) {
            String code = col.getColCode();

            if (i++ > 4) {
                break;
            }

            if (!StringUtils.isEmpty(ret)) {
                ret += ",";
            }
            ret += code;
        }

        userService.saveIssueColumns(ret, user);
        return ret;
    }

}
