import { ComponentBaseMapper, EventStateRegistry, TriggerRenderer } from "@/pages/workflowv2/mappers/types.ts";
import { baseMapper } from "@/pages/workflowv2/mappers/gemini/base.ts";
import { buildActionStateRegistry } from "@/pages/workflowv2/mappers/utils.ts";

export const componentMappers: Record<string, ComponentBaseMapper> = {
  textPrompt: baseMapper,
};

export const triggerRenderers: Record<string, TriggerRenderer> = {};

export const eventStateRegistry: Record<string, EventStateRegistry> = {
  textPrompt: buildActionStateRegistry("completed"),
};
