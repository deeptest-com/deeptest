package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.dao.IssuePageElementDao;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.IssueService;
import com.ngtesting.platform.utils.DateUtil;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class IssueServiceImpl extends BaseServiceImpl implements IssueService {
    Log logger = LogFactory.getLog(IssueServiceImpl.class);

    @Autowired
    IssueDao issueDao;

	@Autowired
	IssuePageDao pageDao;
    @Autowired
    IssuePageElementDao pageElementDao;

    @Override
	public IsuIssue get(Integer id, Integer orgId) {
		IsuIssue po = issueDao.get(id, orgId);

		return po;
	}

	@Override
	public IsuPage getPage(Integer orgId, Integer prjId, String opt) {
        IsuType type = getProjectDefaultType(orgId, prjId);
        Map<String, Integer> pageMap = getProjectDefaultPages(orgId, prjId, type.getId());

		IsuPage page = pageDao.get(pageMap.get(opt), orgId);
		return page;
	}

//    @Override
//    public JSONObject save(JSONObject issue, Integer pageId, TstUser user) {
//        List<IsuPageElement> elems = pageElementDao.listElementByPageId(pageId);
//
//        if (issue.getInteger("id") == null) {
////            issue.setOrgId(user.getDefaultOrgId());
////            issue.setPrjId(user.getDefaultPrjId());
//
//            issueDao.save(issue, elems);
//        } else {
//            Integer count = issueDao.update(issue, elems);
//            if (count == 0) {
//                return null;
//            }
//        }
//
//        return issue;
//    }

    @Override
    public IsuIssue save(JSONObject issue, Integer pageId, TstUser user) {
        List<IsuPageElement> elems = pageElementDao.listElementByPageId(pageId);

        IsuIssue po = null;
        issue.put("orgId", user.getDefaultOrgId());
        issue.put("prjId", user.getDefaultPrjId());

        String uuid = UUID.randomUUID().toString().replace("-", "");
        issue.put("uuid", uuid);

        List<Object> params = genParams(issue, elems, true);
        Integer count = issueDao.save(elems, params);
        if (count > 0) {
            po = issueDao.getByUuid(uuid, user.getDefaultOrgId());
        }

        return po;
    }

    @Override
    public IsuIssue update(JSONObject issue, Integer pageId, TstUser user) {
        List<IsuPageElement> elems = pageElementDao.listElementByPageId(pageId);

        IsuIssue po = null;
        List<Object> params = genParams(issue, elems, false);
        Integer count = issueDao.update(elems, params,
                issue.getInteger("id"), user.getDefaultOrgId());
        if (count > 0) {
            po = issueDao.getByUuid(issue.getString("uuid"), user.getDefaultOrgId());
        }

        return po;
    }

    public List genParams(JSONObject issue, List<IsuPageElement> elems, Boolean init) {
        List<Object> params = new LinkedList<>();

        int i = 0;
        for (IsuPageElement elem : elems) {
            String code = elem.getCode();

            switch(elem.getInput()){
                case "date":
                    params.add(issue.get(code)!=null?issue.getDate(code): null);
                    break;

                default:
                    params.add(issue.get(code));
                    break;
            }
        }

        if (init) {
            params.add(issue.getInteger("orgId"));
            params.add(issue.getInteger("prjId"));
            params.add(issue.getString("uuid"));
        }

        return params;
    }

    @Override
    public IsuType getProjectDefaultType(Integer orgId, Integer prjId) {
        IsuType type = issueDao.getProjectDefaultType(orgId, prjId);
        return type;
    }

    @Override
    public Map<String, Integer>  getProjectDefaultPages(Integer orgId, Integer prjId, Integer typeId) {
        List<Map<String, Object>> ls = issueDao.getProjectDefaultPages(orgId, prjId, typeId);

        Map<String, Integer> ret = new HashMap();
        for (Map<String, Object> map : ls) {
            ret.put(map.get("opt").toString(), Integer.valueOf(map.get("pageId").toString()));
        }
        return ret;
    }

}

