import * as cdk from "aws-cdk-lib";
import {Stack} from "aws-cdk-lib";
import {Role} from "aws-cdk-lib/aws-iam";

export class DummyStack extends Stack {
    role: Role;

    constructor(scope: cdk.App, id: string, props: {
        role?: Role
    }) {
        super(scope, id);
        this.role = new Role(this, 'Role', {
            assumedBy: props.role ?? new cdk.aws_iam.ServicePrincipal('lambda.amazonaws.com'),
        })
        props.role?.grantAssumeRole(this.role);
    }
}
