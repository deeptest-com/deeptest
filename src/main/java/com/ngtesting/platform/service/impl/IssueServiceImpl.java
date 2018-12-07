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
import org.springframework.transaction.annotation.Transactional;

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

    @Override
    @Transactional
    public IsuIssue save(JSONObject issue, Integer pageId, TstUser user) {
        List<IsuPageElement> elemObjs = pageElementDao.listElementByPageId(pageId);
        List<IsuPageElement> elems = genElems(elemObjs);

        IsuIssue po = null;
        issue.put("orgId", user.getDefaultOrgId());
        issue.put("prjId", user.getDefaultPrjId());

        String uuid = UUID.randomUUID().toString().replace("-", "");
        issue.put("uuid", uuid);

        List<Object> params = genParams(issue, elems, true);

        List<IsuPageElement> elems1 = new LinkedList<>();
        List<Object> params1 = new LinkedList<>();
        List<IsuPageElement> elems2 = new LinkedList<>();
        List<Object> params2 = new LinkedList<>();

        genDataForExtTable(elems, params, elems1, params1, elems2, params2);

        Integer count = issueDao.save(elems1, params1);
        if (count > 0) {
            po = issueDao.getByUuid(uuid, user.getDefaultOrgId());
            count = issueDao.saveExt(elems2, params2, po.getId());
        }

        if (count > 0) {
            po = issueDao.get(po.getId(), user.getDefaultOrgId());
        }

        return po;
    }

    @Override
    @Transactional
    public IsuIssue update(JSONObject issue, Integer pageId, TstUser user) {
        List<IsuPageElement> elems = pageElementDao.listElementByPageId(pageId);

        IsuIssue po = null;
        List<Object> params = genParams(issue, elems, false);

        List<IsuPageElement> elems1 = new LinkedList<>();
        List<Object> params1 = new LinkedList<>();
        List<IsuPageElement> elems2 = new LinkedList<>();
        List<Object> params2 = new LinkedList<>();

        genDataForExtTable(elems, params, elems1, params1, elems2, params2);

        Integer count = issueDao.update(elems1, params1, issue.getInteger("id"), user.getDefaultOrgId());
        if (count > 0 && params2.size() > 0) {
            count = issueDao.updateExt(elems2, params2, issue.getInteger("id"));
        }

        if (count > 0) {
            po = issueDao.get(issue.getInteger("id"), user.getDefaultOrgId());
        }

        return po;
    }

    private List<IsuPageElement> genElems(List<IsuPageElement> elemObjs) {
        List<IsuPageElement> elems = new LinkedList<>();
        for (IsuPageElement elem : elemObjs) {
            elems.add(elem);
        }

        elems.add(new IsuPageElement("orgId", ""));
        elems.add(new IsuPageElement("projectId", ""));
        elems.add(new IsuPageElement("uuid", ""));

        return elems;
    }

    public List genParams(JSONObject issue, List<IsuPageElement> elems, Boolean init) {
        List<Object> params = new LinkedList<>();

        int i = 0;
        for (IsuPageElement elem : elems) {
            String code = elem.getColCode();

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

    @Override
    public void genDataForExtTable(List<IsuPageElement> elems,
                                   List<Object> params,
                                   List<IsuPageElement> elems1,
                                   List<Object> params1,
                                   List<IsuPageElement> elems2,
                                   List<Object> params2) {
        int i = 0;
        for (IsuPageElement elem: elems) {
            if (!elem.getColCode().startsWith("prop")) {
                elems1.add(elem);
                params1.add(params.get(i));
            } else {
                elems2.add(elem);
                params2.add(params.get(i));
            }

            i++;
        }
    }

}

