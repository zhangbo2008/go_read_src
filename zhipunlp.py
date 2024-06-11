# -*- coding:utf-8 -*-
# pip install zhipuai 请先在终端进行安装 #这个脚本用来我读代码用.

from zhipuai import ZhipuAI

client = ZhipuAI(api_key="a3c9f5bb3b234522807f91c1e4e420a9.CgmSKoI5C1COL9TX")


import time
aaaaa=time.time()
response = client.chat.completions.create(
model="glm-4-0520",
    messages=[
        {
            "role": "system",
            "content": "你是一位非常专业的代码专家。请讲解下面代码" 
        },
        {
            "role": "user",
            "content": '''
               // uint64 atomicload64(uint64 volatile* addr);
// so actually
// void atomicload64(uint64 *res, uint64 volatile *addr);
TEXT runtime·atomicload64(SB), NOSPLIT, $0-12
	MOVL	ptr+0(FP), AX
	TESTL	$7, AX
	JZ	2(PC)
	MOVL	0, AX // crash with nil ptr deref
	LEAL	ret_lo+4(FP), BX
	// MOVQ (%EAX), %MM0
	BYTE $0x0f; BYTE $0x6f; BYTE $0x00
	// MOVQ %MM0, 0(%EBX)
	BYTE $0x0f; BYTE $0x7f; BYTE $0x03
	// EMMS
	BYTE $0x0F; BYTE $0x77
	RET

                '''
        }
    ],
    top_p= 1,
    temperature= 0.95,
    max_tokens=1024,
    tools = [{"type":"web_search","web_search":{"search_result":False}}],
    stream=False,
)
import json
dict()
a=response.choices[0].message.content
print(a)


print('使用的时间',time.time()-aaaaa)
