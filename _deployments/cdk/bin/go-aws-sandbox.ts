#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { GoAwsSandboxStack } from '../lib/go-aws-sandbox-stack';

const app = new cdk.App();
new GoAwsSandboxStack(app, 'GoAwsSandboxStack', {
  env: { account: process.env.CDK_DEFAULT_ACCOUNT, region: process.env.CDK_DEFAULT_REGION },
});