FROM ubuntu

# 全局工作目录
WORKDIR /workspace/leyanapi-backend

COPY application /workspace/leyanapi-backend/
COPY config/ /workspace/leyanapi-backend/config/


EXPOSE 8085 8086

ENTRYPOINT ["/bin/sh", "-c"]  
CMD ["cp /workspace/leyanapi-backend/config/configMap/*.yaml /workspace/leyanapi-backend && chmod +x /workspace/leyanapi-backend/application && /workspace/leyanapi-backend/application"]
