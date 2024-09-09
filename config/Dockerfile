FROM ubuntu

# 全局工作目录
WORKDIR /workspace/thirdpartyapi-backend

COPY application /workspace/thirdpartyapi-backend/
COPY config/ /workspace/thirdpartyapi-backend/config/


EXPOSE 8085 8086

ENTRYPOINT ["/bin/sh", "-c"]
CMD ["cp /workspace/thirdpartyapi-backend/config/configMap/*.yaml /workspace/thirdpartyapi-backend && chmod +x /workspace/thirdpartyapi-backend/application && /workspace/thirdpartyapi-backend/application"]
