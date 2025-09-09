---
id: AIOPS
aliases: []
tags: []
created: 2025-06-12 18:57:32
modified: 2025-11-11 18:56:45
---


[[Terminologies]]



172.16.10.47
10.100.58.1

- [~] traceroute windows host port 
- [x] snpm v3 command
- [x] icmp
- [x] netstats
- [x] dhcp

fping hops
port state
time wait state
closed closed wait
sync finwait
winrm


ZMQ
pull and multi push
io workers
timeout 

vertx context



| Model (open)                                                         |            Typical sizes |                               Rough VRAM to run (inference) | Recommended GPUs (local / small server / pro)                                               |                                                       Example GPU price (new, late-2025 est.) |
| -------------------------------------------------------------------- | -----------------------: | ----------------------------------------------------------: | ------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------: |
| **Mistral Mixtral / Mixtral-8x** (Mixtral 8x7B / Mixtral 8x22B etc.) | 7B / 22B (SMoE variants) | **~12–24 GB** (7B quantized) / 24–48 GB for larger variants | Local: RTX 4090 (24GB); Server: 1–2×4090 or 1×RTX 5090; Pro: AMD MI250 / NVIDIA H100 (80GB) |                         RTX 4090 ~ $1.6k–$3k street (varies); H100 ~ $25k+. ([Mistral AI][1]) |
| **Llama 3 (Meta)** — 8B / 70B                                        |                 8B / 70B |   8B: **~10–16 GB**; 70B: **~40–80+ GB** (depends on quant) | Local: 4090 for 8B; Server: 2×4090 or A100/MI300 for 70B; Pro: H100 80GB                    |               Llama 3 announced publicly; fits smaller GPUs with quantization. ([Meta AI][2]) |
| **Falcon 180B** (TII / Falcon family)                                |                     180B | **~120–200+ GB** native — needs sharding/SMoE/quant+offload | Server/professional: multi-GPU (H100 / A100 / MI300) or CPU+GPU clusters                    | Falcon available as open weights; typically requires datacenter GPUs. ([falconllm.tii.ae][3]) |
| **Mosaic MPT / MPT-30B**                                             |                   7B–30B |                                          30B: **~40–80 GB** | Local: 4090 for smaller; server: 2–4×4090 / A100 / MI250                                    |                                           MPT family widely used for local setups with quant. |
| **RedPajama / RWKV / other <=13B**                                   |                   3B–13B |                                                **~6–16 GB** | Local: RTX 4070/4070 Ti / 4090                                                              |                                                                Good for single-GPU local use. |
| **Gemma / Open models by other labs** (varies)                       |                   3B–70B |                                                      varies | similar to Llama / Mistral recommendations                                                  |                                                                                               |

[1]: https://mistral.ai/news/mixtral-of-experts?utm_source=chatgpt.com "Mixtral of experts"
[2]: https://ai.meta.com/blog/meta-llama-3/?utm_source=chatgpt.com "Introducing Meta Llama 3: The most capable openly ..."
[3]: https://falconllm.tii.ae/falcon-180b.html?utm_source=chatgpt.com "Falcon 180B"
