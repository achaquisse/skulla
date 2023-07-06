import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import {Nextjs} from "cdk-nextjs-standalone";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class FrontendInfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    new Nextjs(this, 'skulla-ui', {
      nextjsPath: '../frontend-ui',
      buildCommand: 'npx --yes open-next@1.4.0 build'
    })
  }
}
