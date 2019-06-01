# AWSPrice CLI Tool

### Usage

#### For Linux 
```
$wget https://github.com/aws-user-group-myanmar-aws-ugm/awsprice/releases/download/v1.0.0/awsprice
$chmod +x awsprice
```

#### For Window 
```
$wget https://github.com/aws-user-group-myanmar-aws-ugm/awsprice/releases/download/v1.0.0/awsprice.exe
```

#### For MAC
```
$wget https://github.com/aws-user-group-myanmar-aws-ugm/awsprice/releases/download/v1.0.0/awsprice.mac
$chmod +x awsprice.mac
```


```shell
$./awsprice help
Aws Price list tools

Usage:
  awsprice [command]

Available Commands:
  get         get <AWS resource>
  help        Help about any command
  version     awsprice version

Flags:
  -h, --help   help for awsprice

Use "awsprice [command] --help" for more information about a command.
```

```
$./awsprice get
Error: You must specify the type of resource to get
Usage:
  awsprice get [flags]
  awsprice get [command]

Available Commands:
  appstream      A brief price description of AWS App Stream
  athena         A brief price description of AWS Athena
  backup         A brief price description of AWS Back up
  certmgr        A brief price description of AWS Certificate Manager
  cloudfront     A brief price description of AWS Cloud Front
  cloudmap       A brief price description of AWS Cloud Map
  cloudsearch    A brief price description of AWS Cloud Search
  cloudwatch     A brief price description of AWS Cloud Watch
  codebuild      A brief price description of AWS Cloud Build
  codecommit     A brief price description of AWS Code Commit
  codedeploy     A brief price description of AWS Code Deploy
  codepipeline   A brief price description of AWS Code Pipline
  cognito        A brief price description of AWS Cognito
  datapipeline   A brief price description of AWS Data Pipeline
  devsupport     A brief price description of AWS Developer Support
  dirservice     A brief price description of AWS Directory Service
  dynamodb       A brief price description of AWS Dynamodb
  ec2            A brief price description of AWS EC2 instances
  ec3            A brief description of your command
  ecr            A brief price description of AWS Container Registry
  efs            A brief price description of AWS EFS
  eks            A brief price description of Aws Kubernetes Service
  escache        A brief price description of AWS Elastic Cache
  esmapreduce    A brief price description of AWS Elasltic Map Reduce
  gamelift       A brief price description of AWS Game Lift
  iot            A brief price description of IOT
  lambda         A brief price description of AWS Lambda
  mgmtblockchain A brief price description of AWS Manage BlockChain
  ml             A brief price description of Machine Learning
  opsworks       A brief price description of AWS OpsWorks
  queueservice   A brief price description of AWS Queue Service
  rds            A brief price description of Amazon RDS
  redshift       A brief price description of Amazon Redshift
  route53        A brief price description of AWS Route53
  s3             A brief price description of Amazon S3
  secretmanager  A brief price description of AWS Secret Manager
  simpledb       A brief price description of Amazon Simple DB
  sns            A brief price description of AWS SNS
  vpc            A brief price description of AWS VPC
  workspaces     A brief price description of AWS WorkSpaces

Flags:
  -h, --help   help for get

Use "awsprice get [command] --help" for more information about a command.

You must specify the type of resource to get
```

```
$./awsprice get rds
+--------------------------------+----------------+-------------+
|          DESCRIPTION           |      USD       |    UNIT     |
+--------------------------------+----------------+-------------+
| $0.052 per RDS db.t1.micro     |   0.0520000000 | Hrs         |
| Multi-AZ instance hour         |                |             |
| (or partial hour) running      |                |             |
| PostgreSQL                     |                |             |
+--------------------------------+----------------+-------------+
| $0.646 per RDS db.m4.xlarge    |   0.6460000000 | Hrs         |
| instance hour (or partial      |                |             |
| hour) running SQL Server EE    |                |             |
| (BYOL)                         |                |             |
+--------------------------------+----------------+-------------+
| USD 1.5936 per db.t2.2xlarge   |   1.5936000000 | Hrs         |
| Multi-AZ instance hour (or     |                |             |
| partial hour) running Oracle   |                |             |
| SE (BYOL)                      |                |             |
+--------------------------------+----------------+-------------+
| Oracle SE (BYOL),              |   0.6080000000 | Hrs         |
| db.t2.2xlarge reserved         |                |             |
| instance applied               |                |             |
+--------------------------------+----------------+-------------+
| Oracle SE (BYOL),              |   0.6080000000 | Hrs         |
| db.t2.2xlarge reserved         |                |             |
| instance applied               |                |             |
+--------------------------------+----------------+-------------+
```

```
$./awsprice get ec2 --os Linux --tenancy Shared --type t2.micro
+--------------------------------+--------------+
|          DESCRIPTION           |     USD      |
+--------------------------------+--------------+
| $0.0134 per On Demand Linux    | 0.0134000000 |
| t2.micro Instance Hour         |              |
+--------------------------------+--------------+
| $0.0126 per On Demand Linux    | 0.0126000000 |
| t2.micro Instance Hour         |              |
+--------------------------------+--------------+
| $0.0862 per On Demand Linux    | 0.0862000000 |
| with SQL Web t2.micro Instance |              |
| Hour                           |              |
+--------------------------------+--------------+
| $0.0146 per On Demand Linux    | 0.0146000000 |
| t2.micro Instance Hour         |              |
+--------------------------------+--------------+
| $0.0144 per On Demand Linux    | 0.0144000000 |
| t2.micro Instance Hour         |              |
+--------------------------------+--------------+
| $0.0146 per On Demand Linux    | 0.0146000000 |
| t2.micro Instance Hour         |              |
+--------------------------------+--------------+
```


#### Note: EC2 and Some Will take time between 15 min and 20 min depend on your connection.
