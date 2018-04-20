package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.config.PropertyConfig;
import com.ngtesting.platform.service.JenkinsService;
import com.ngtesting.platform.vo.*;
import com.offbytwo.jenkins.JenkinsServer;
import com.offbytwo.jenkins.model.Job;
import com.offbytwo.jenkins.model.JobWithDetails;
import org.springframework.stereotype.Service;

import java.net.URI;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static com.ngtesting.platform.config.Constant.JenkinsTask;

@Service
public class JenkinsServiceImpl extends BaseServiceImpl implements JenkinsService {

    @Override
    public AiRunVo genRunVo(AiTestTaskVo task) {
        AiRunVo run = new AiRunVo();
        AiRunEnvi env = new AiRunEnvi();
        AiRunRes res = new AiRunRes();
        List<AiRunMlf> mlfs = JSON.parseObject(task.getMlfs(), List.class);
        res.setMlfs(mlfs);
        run.setEnvi(env);
        run.setRes(res);

        env.setTestType(task.getTestType());
        env.setSite(task.getTestEnv());
        env.setResParam(task.getAsrLangModel());
        env.setProductId(task.getTestProductId().toString());
        env.setStartIndex(task.getStartIndex());
        env.setNumbToRun(task.getNumbToRun());
        env.setAudioType(task.getAudioType());
        env.setAliasKey(task.getProductBranch());
        env.setIsFuse(task.getFuse().toString());

        return run;
    }

    @Override
    public String execute(AiTestTaskVo vo) {
        AiRunVo runVo = genRunVo(vo);
        String json = JSON.toJSONString(runVo);

        JenkinsServer jenkinsServer;
        try {
            String jenkinsSvr = PropertyConfig.getConfig("jenkins.server");
            String jenkinsUser = PropertyConfig.getConfig("jenkins.user");
            String jenkinsPassword = PropertyConfig.getConfig("jenkins.password");

            jenkinsServer = new JenkinsServer(new URI(jenkinsSvr), jenkinsUser, jenkinsPassword);
            Map<String, Job> jobs = jenkinsServer.getJobs();
            JobWithDetails job = jobs.get(JenkinsTask.get(vo.getTestType())).details();

            Map<String, String> params = new HashMap();
            params.put("json", JSON.toJSONString(runVo));
            params.put("suite", JenkinsTask.get(vo.getTestType()));

            job.build(params, true);

        } catch (Exception e) {
            e.printStackTrace();
        }

        return json;
    }
}

