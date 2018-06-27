package com.ngtesting.platform.utils;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ShellUtil {
    private static Logger logger = LoggerFactory.getLogger(ShellUtil.class);

    public static final int DEFAULT_TIMEOUT = 30 * 1000; // in millseconds
    public static final int DEFAULT_INTERVAL = 1000;

    /**
     * Executes the specified command in a separate process. The method then
     * blocks until the process returned. If an error arises during the
     * execution or if the exeecuted process returned an non-null return code,
     * the content of the process' stderr is returned to the caller. If the
     * execution is fine, null is returned.
     *
     * @param command String
     * @return CommandResult
     */
    public static ShellResult exec(String cmd, Map<String, String> regex) {
        long start = System.currentTimeMillis();
        long last;

        ShellResult commandResult = new ShellResult();

        try {
            Process process = Runtime.getRuntime().exec(new String[]{"/bin/sh", "-c", cmd});
            process(process, commandResult, regex);

            if (process != null) {
                process.destroy();
            }

            last = (System.currentTimeMillis() - start) / 1000;
            logger.info("Execute command [" + cmd + "], last [" + last
                    + "] s.");

        } catch (Exception e) {
            last = (System.currentTimeMillis() - start) / 1000;
            String error = "Execute command [" + cmd + "] last [" + last
                    + "] s, failed [" + e.getMessage() + "]";
            logger.error(error, e);

            commandResult.setExitValue(ShellResult.COMMAND_EXIT_VALUE_EXCEPTION);
            commandResult.setErrorOutput(error);
        }

        return commandResult;
    }

    private static void process(Process process, ShellResult res, Map<String, String> regex) {
        BufferedReader errorReader = null;
        BufferedReader inputReader = null;

        try {
            errorReader = new BufferedReader(new InputStreamReader(
                    process.getErrorStream()));
            inputReader = new BufferedReader(new InputStreamReader(
                    process.getInputStream()));

            // timeout control
            long start = System.currentTimeMillis();
            boolean processFinished = false;

            while (System.currentTimeMillis() - start < DEFAULT_TIMEOUT
                    && !processFinished) {
                processFinished = true;
                try {
                    process.exitValue();
                } catch (IllegalThreadStateException e) {
                    // process hasn't finished yet
                    processFinished = false;

                    try {
                        Thread.sleep(DEFAULT_INTERVAL);
                    } catch (InterruptedException e1) {
                        logger.error(
                                "Process, failed [" + e.getMessage() + "]", e);
                    }
                }
            }

            if (!processFinished) {
                res.setExitValue(ShellResult.COMMAND_EXIT_VALUE_TIMEOUT);
                res.setErrorOutput("远程手机操作超时");
                return;
            }

            res.setExitValue(process.waitFor());

            StringBuffer sb;
            String line;

            // parse error info
            if (errorReader.ready()) {
                sb = new StringBuffer();
                while ((line = errorReader.readLine()) != null) {
                    sb.append(line);
                    System.out.println(line);
                }
                res.setErrorOutput(sb.toString());
            }

            // parse info
            if (inputReader.ready()) {
                sb = new StringBuffer();
                while ((line = inputReader.readLine()) != null) {
                    sb.append(line);
                    System.out.println(line);

                    if (regex != null && regex.keySet().size() > 0) {
                        for (String key : regex.keySet()) {
                            String r = regex.get(key);
                            Pattern pattern = Pattern.compile(r);

                            Matcher matcher = pattern.matcher(line);
                            if (matcher.find()) {
                                String s = matcher.group(1);
                                if (res.getRegexMap().containsKey(key)) {
                                    res.getRegexMap().put(key,
                                            res.getRegexMap().get(key) + "," + s); // 多行命中，
                                    // 返回逗号风格的字符串
                                } else {
                                    res.getRegexMap().put(key, s); // 每行中字返回第一个匹配的值
                                }
                            }
                        }
                    }
                    if (line.indexOf("ERROR") > -1 || line.indexOf("Error") > -1
                            || line.indexOf("error") > -1) {
                        res.setExitValue(ShellResult.COMMAND_EXIT_VALUE_ERROR_IN_OUTPUT);
                        res.setErrorOutput("find 'ERROR' in command output");
                    }
                }
                res.setInfoOutput(sb.toString());
            }

        } catch (Exception e) {
            String error = "Command process, failed [" + e.getMessage() + "]";
            logger.error(error, e);

            res.setExitValue(ShellResult.COMMAND_EXIT_VALUE_EXCEPTION);
            res.setErrorOutput(error);

        } finally {
            if (errorReader != null) {
                try {
                    errorReader.close();
                } catch (IOException e) {
                    logger.error(
                            "Close BufferedReader, failed [" + e.getMessage()
                                    + "]", e);
                }
            }

            if (inputReader != null) {
                try {
                    inputReader.close();
                } catch (IOException e) {
                    logger.error(
                            "Close BufferedReader, failed [" + e.getMessage()
                                    + "]", e);
                }
            }
        }
    }

}
