import * as cdk from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { GoAwsSandboxStack as TestStack } from '../lib/go-aws-sandbox-stack';
import * as fs from 'fs';
import * as path from 'path';

test('GoEcsBatchSampleStack Snapshot Test', () => {
    const env = "stage"
    const cdkJson = JSON.parse(fs.readFileSync(path.resolve(__dirname, "../cdk.json"), "utf8"))
    const context = Object.assign({ env }, cdkJson?.context)
    const app = new cdk.App({ context });
    // WHEN
    const stack = new TestStack(app, 'TestStack');
    // THEN
    const template = Template.fromStack(stack);

    expect(template.toJSON()).toMatchSnapshot();
});
