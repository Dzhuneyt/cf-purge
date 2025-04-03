#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import {DummyStack} from "../lib/DummyStack";

const app = new cdk.App();

for (let i = 0; i < 3; i++) {
    const producerStack = new DummyStack(app, `cf-purge-test-CONSUMER-STACK-LEVEL-${i}`, {});
    new DummyStack(app, `cf-purge-test-PRODUCER-STACK-LEVEL-${i}`, {
        role: producerStack.role,
    });
}
