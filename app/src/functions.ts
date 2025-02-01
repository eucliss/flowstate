import type { Node, Edge } from '@vue-flow/core'  
import { ref } from 'vue'
export const URL = 'http://localhost:3000'

// SELECT * FROM 'test_torq3'

export const nodeTypes = ['default', 'custom']

export type ComparisonType = {
    leftValue: string
    rightValue: string
    operator: string
}

export type NodeData = {
    label: string
    sql: string
    successRoute: ComparisonType
    failureRoute: ComparisonType
    status: boolean
}

export type EdgeData = {
    id: string
    source: string
    target: string
    animated: boolean
}

export const nodes = ref<Node[]>([])
export const edges = ref<Edge[]>([])

export const convertNodetoGoNode = (node: Node) => {
    console.log("Converting node to Go node: ", node)
    return {
        id: node.id,
        type: node.type,
        x: node.position.x,
        y: node.position.y,
        label: node.data.label,
        sql: node.data.sql,
        successRoute: {
            leftValue: node.data.successRoute.leftValue,
            rightValue: node.data.successRoute.rightValue,
            operator: node.data.successRoute.operator,
        },
        failureRoute: {
            leftValue: node.data.failureRoute.leftValue,
            rightValue: node.data.failureRoute.rightValue,
            operator: node.data.failureRoute.operator,
        },
    }
}

export const convertGoNodeToNode = (n: { 
    Id: string, 
    Type: string, 
    X: number, 
    Y: number, 
    Label: string, 
    SQL: string,
    SuccessRoute: ComparisonType,
    FailureRoute: ComparisonType,
}) => {
    console.log("Converting Go node to node: ", n)
    return {
        id: n["id"],
        type: n["type"],
        position: {
            x: n["x"],
            y: n["y"],
        },
        data: {
            label: n["label"],
            sql: n["sql"],
            successRoute: n["successRoute"],
            failureRoute: n["failureRoute"],
        },
    }
}

export const convertGoEdgeToEdge = (e: { Id: string, Source: string, Target: string, Animated: boolean }) => {
    console.log("Converting Go edge to edge: ", e)
    return {
        id: e["Id"],
        source: e["Source"],
        target: e["Target"],
        animated: e["Animated"],
    }
}

export const convertEdgeToGoEdge = (e: Edge) => {
    console.log("Converting edge to Go edge: ", e)
    return {
        Id: e.id,
        Source: e.source,
        Target: e.target,
        Animated: e.animated,
    }
}

export const getNodeStatus = async (node: Node) => {
    const response = await fetch(`${URL}/node-route-status`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            id: node.id,
            route: "success",
        }),
    })
    const data = await response.json()
    console.log("Node status: ", data)
    return data
}

export const updateAllNodesStatus = async () => {
    nodes.value.forEach(async (node) => {
        if (node.type === "custom") {
            const status = await getNodeStatus(node)
            node.data.status = status
        }
    })
    console.log("All nodes status updated: ", nodes.value)
}

export const getFlowState = async () => {
    console.log("Loading flow state ... ")
    const response = await fetch(`${URL}/load-flow-state`, {
        method: 'GET',
    })
    const data = await response.json()

    nodes.value = data.nodes.map(convertGoNodeToNode)
    edges.value = data.edges.map(convertGoEdgeToEdge)
    await updateAllNodesStatus()
    console.log("Flow state loaded: ", nodes)
}

export const updateNode = async (node: Node) => {
    console.log("Updating node: ", node)
    nodes.value = nodes.value.map(n => 
        n.id === node.id ? { ...node } : n
    )
    const response = await fetch(`${URL}/update-node`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(convertNodetoGoNode(node)),
    })
    const data = await response.json()
    return data
}

export const addNode = async (node: Node) => {
    const goNode = convertNodetoGoNode(node)
    nodes.value.push(node)
    console.log("Adding node to Go: ", goNode)
    const response = await fetch(`${URL}/add-node`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(goNode),
    })
    const data = await response.json()
    return data
}

export const connectEdge = async (edge: Edge) => {
    const response = await fetch(`${URL}/connect-edge`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(edge),
    })
    const data = await response.json()
    return data
}

export const resetFlowState = async () => {
    const response = await fetch(`${URL}/reset`, {
        method: 'GET',
    })
    const data = await response.json()
    return data
}

export const testing = () => {
    console.log('testing')
}