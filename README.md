# auto_block_snapshot

auto_block_snapshot is a tool designed to automate the process of snapshotting, pruning, and uploading data to an AWS S3 bucket from a blockchain full node built with Cosmos-SDK and Tendermint.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)

## Installation
As a prerequisite, you'll need a Cosmos fulnode and a Cosmo pruner binary.
```
https://github.com/binaryholdings/cosmprund
https://github.com/cosmos/gaia 
```
You can also set up AWS environment variables or credentials as shown below.

### Environment Variables:

```
export AWS_ACCESS_KEY_ID=your_access_key
export AWS_SECRET_ACCESS_KEY=your_secret_key
export AWS_REGION=your_aws_region
```
### AWS Credentials File:
Alternatively, you can create the file ~/.aws/credentials (on Linux, macOS, or Unix) or %UserProfile%\.aws\credentials (on Windows). This file should contain lines in the following format:

```
[default]
aws_access_key_id = your_access_key
aws_secret_access_key = your_secret_key
region = your_aws_region
```


This project is written in Go. Make sure you have Go installed on your machine.

```
git clone https://github.com/b-harvest/auto_block_snapshot.git
cd auto_block_snapshot
go build .
./auto_block_snapshot
or
./auto_block_snapshot -c /path/to/config.toml
```


## Configuration
The configuration of this tool can be done using a `config.toml` file. An example configuration can be found in `config.toml`.

The `config.toml` should contain the following fields:

- `fullnode.path`: path to the full node binary
- `fullnode.data_path`: path to the chain data Folder
- `fullnode.chain_name`: The name of the blockchain
- `pruner.path`: path to the cosmopruner binary
- `aws.region`: AWS region
- `aws.bucket`: AWS S3 bucket name

## set up an AWS account to use S3.
### 1. Create an AWS Account

If you do not have an AWS account, you need to create one. Go to the AWS homepage and click on the "Create an AWS Account" button. Follow the online instructions. Part of the sign-up procedure involves receiving a phone call and entering a verification code on the phone keypad.

### 2. Sign into the AWS Management Console

Once you have an AWS account, sign in to the AWS Management Console.

### 3. Create an S3 Bucket

In the AWS Management Console, type S3 in the service search bar and click on the S3 tile to open the S3 dashboard. Click "Create Bucket". Enter a DNS-compliant name for your bucket and select a region. Leave the rest of the options at their defaults for now, and click "Create".

### 4. Configure Bucket Permissions

- Block All Public Access: By default, all public access is blocked for new buckets. If you want your bucket to be publicly accessible, you can uncheck this option. However, be aware of the potential security implications.

- Bucket Policy: This is a JSON file that grants or denies permissions to your bucket. If you want your application to access the bucket, you will need to grant it permissions using a bucket policy.

### 5. Create IAM User

You will need to create an IAM user that your application will use to access the bucket. In the AWS Management Console, go to the IAM (Identity and Access Management) service. Click on "Users" in the left-hand menu and then click on "Add User". Enter a user name, select "Programmatic access" and click "Next".

- Set Permissions: On the permissions screen, select "Attach existing policies directly". In the policy filter, type "S3" and check the box next to "AmazonS3FullAccess". Click "Next" until you get to the review screen, then click "Create User".

- Store Credentials: After creating the user, you will see a screen that shows the user's Access Key ID and Secret Access Key. Store these credentials securely! Your application will use these credentials to access the S3 bucket. The Secret Access Key will only be shown this once, so make sure to save it.

### More
Your AWS S3 bucket is now ready to be used with your application. In your config.toml file, set the aws.region to the region you chose when creating the bucket, and aws.bucket to the name of the bucket. Your application will use the IAM user's Access Key ID and Secret Access Key to authenticate with AWS and gain access to the bucket. These can be set in your environment variables or in your AWS credentials file.