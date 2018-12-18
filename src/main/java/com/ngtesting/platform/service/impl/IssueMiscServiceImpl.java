package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueMiscDao;
import com.ngtesting.platform.model.IsuLinkReason;
import com.ngtesting.platform.model.IsuTag;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueMiscService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class IssueMiscServiceImpl extends BaseServiceImpl implements IssueMiscService {
    @Autowired
    IssueMiscDao issueMiscDao;

}

