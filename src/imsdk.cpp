#include "imsdk.h"

#include <iostream>
#include <sstream>
#include "json.hpp"
using json = nlohmann::json;

char *GetOperationId(const char *prefix)
{
    static int operationIndex = 0;
    operationIndex = operationIndex + 1;
    std::stringstream ss;
    ss << prefix << "_" << operationIndex;
    return (char *)ss.str().c_str();
}

void OpenIM::InitSDK()
{
    init_sdk(GetOperationId("init_sdk"), "");
}