<script setup lang="ts">
import type { NodeProps } from '@vue-flow/core'
import { NodeResizer } from '@vue-flow/node-resizer'
import { ref, onMounted, computed } from 'vue'

const props = defineProps<NodeProps & {
    updateNodeData?: (nodeId: string, newData: any) => void
}>()
const emit = defineEmits(['updateNodeInternals', 'updateNodeContent'])
const isEditing = ref(false)
const label = computed(() => props.data.label || "Text Node")

const PLACEHOLDER_TEXT = "Type your text here..."

const content = computed({
    get: () => {
        // Return empty string if the content is the placeholder
        return props.data.content === PLACEHOLDER_TEXT ? "" : (props.data.content || PLACEHOLDER_TEXT)
    },
    set: (newValue) => {
        if (props.updateNodeData) {
            // If the new value is empty, set it to the placeholder
            const valueToSave = newValue.trim() === "" ? PLACEHOLDER_TEXT : newValue
            props.updateNodeData(props.id, {
                ...props.data,
                content: valueToSave
            })
        }
    }
})

// Add a computed property to handle display text
const displayContent = computed(() => {
    return props.data.content || PLACEHOLDER_TEXT
})

const toggleEdit = () => {
    if (isEditing.value) {
        return
    }
    isEditing.value = !isEditing.value
}

const handleBlur = (event: FocusEvent) => {
    // Check if we clicked the save button to prevent double-saving
    if (!(event.relatedTarget as HTMLElement)?.classList.contains('save-btn')) {
        isEditing.value = false
        if (props.updateNodeData) {
            props.updateNodeData(props.id, {
                ...props.data,
                content: content.value
            })
        }
    }
}

const handleSave = () => {
    isEditing.value = false  // Set this first
    if (props.updateNodeData) {
        props.updateNodeData(props.id, {
            ...props.data,
            content: content.value
        })
    }
}

</script>

<template>
    <NodeResizer :minWidth="200" :minHeight="50" />
    <div class="text-node-container vue-flow__node-textNode">
        <div class="node-header">{{ label }}</div>
        <div class="content-wrapper" @click="toggleEdit">
            <div v-if="isEditing" class="edit-container">
                <textarea class="nodrag text-input" 
                         v-model="content"
                         @blur="handleBlur"
                         ref="textArea"
                         autofocus>
                </textarea>
                <!-- <button class="save-btn nodrag" 
                        @click="handleSave"
                        @mousedown.prevent>
                    Save
                </button> -->
            </div>
            <div v-else 
                 class="html-content"
                 v-html="displayContent">
            </div>
        </div>
    </div>
</template>

<style scoped>
@import '../../assets/nodes.css';
@import 'https://cdn.jsdelivr.net/npm/@vue-flow/core@1.42.1/dist/style.css';
@import 'https://cdn.jsdelivr.net/npm/@vue-flow/core@1.42.1/dist/theme-default.css';
@import 'https://cdn.jsdelivr.net/npm/@vue-flow/controls@latest/dist/style.css';
@import 'https://cdn.jsdelivr.net/npm/@vue-flow/minimap@latest/dist/style.css';
@import 'https://cdn.jsdelivr.net/npm/@vue-flow/node-resizer@latest/dist/style.css';

.text-node-container {
    min-width: 200px;
    min-height: 50px;
    width: 100%;
    height: 100%;
    padding: 0;
    border-radius: 4px;
    display: flex;
    flex-direction: column;  /* Change back to column for title on top */
    background: white;  /* Change to white background */
}

.node-header {
    padding: 4px 8px;
    background: white;
    font-size: 12px;
    font-weight: bold;
    color: #666;
    text-align: left;
    border-radius: 4px 4px 0 0;  /* Round top corners */
}

.content-wrapper {
    flex: 1;
    padding: 8px;
    display: flex;
    cursor: pointer;
    background: white;  /* Ensure content area is also white */
}

.edit-container {
    position: relative;
    width: 100%;
    height: 100%;
}

.save-btn {
    position: absolute;
    bottom: 8px;
    right: 8px;
    padding: 6px 12px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 12px;
    font-weight: 500;
    transition: background-color 0.2s;
    z-index: 10;
}

.save-btn:hover {
    background-color: #45a049;
}

.text-input {
    width: 100%;
    height: 100%;
    padding-bottom: 40px; /* Make room for the save button */
    border: none;
    padding: 4px 8px;
    border-radius: 2px;
    background: #ffffff;
    font-size: 14px;
    color: #000000;  /* Explicit text color */
    resize: none;  /* Prevents textarea resize handle */
    font-family: inherit;  /* Maintains consistent font */
}

.text-input:focus {
    outline: none;
    box-shadow: 0 0 0 2px rgba(0, 89, 220, 0.3);
}

.html-content {
    width: 100%;
    height: 100%;
    padding: 4px 8px;
    overflow-y: auto;
    color: #000000;
    cursor: pointer;
}

/* Add styles for placeholder text */
.html-content:empty::before,
.html-content:has(:empty)::before {
    content: attr(data-placeholder);
    color: #999;
    font-style: italic;
}

/* Ensure HTML content inherits text styles */
.html-content :deep(*) {
    font-family: inherit;
    font-size: inherit;
    color: #000000;  /* Ensure all nested elements have black text */
}
</style>
