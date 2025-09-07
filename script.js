import pubsub from "k6/x/pubsub";
import { SharedArray } from "k6/data";

// Create Pub/Sub client using Config object
const client = pubsub.publisher({
    project_id: "flink-core-shared",
    topic: "minimal-event-schemaless"
});
if (!client) console.error("Failed to create client");

// Load test data once
const testData = new SharedArray('test data', function() {
  return JSON.parse(open('./testdata.json'));
});

// Convert each object to a JSON string
const messages = testData.map(item => JSON.stringify(item));

export default function () {
    // Publish all messages in batch
    const ids = pubsub.publishBatch(client, messages);
    console.log("Published messages successfully:", ids);
}
