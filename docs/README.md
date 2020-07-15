# Prado::Rundeck::Project

Rundeck Project Resource

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Prado::Rundeck::Project",
    "Properties" : {
        "<a href="#rundeckauthtoken" title="RundeckAuthToken">RundeckAuthToken</a>" : <i>String</i>,
        "<a href="#rundeckapiurl" title="RundeckAPIURL">RundeckAPIURL</a>" : <i>String</i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectID">ProjectID</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Prado::Rundeck::Project
Properties:
    <a href="#rundeckauthtoken" title="RundeckAuthToken">RundeckAuthToken</a>: <i>String</i>
    <a href="#rundeckapiurl" title="RundeckAPIURL">RundeckAPIURL</a>: <i>String</i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#projectid" title="ProjectID">ProjectID</a>: <i>String</i>
</pre>

## Properties

#### RundeckAuthToken

The Rundeck API Auth Token. See more https://docs.rundeck.com/docs/api/rundeck-api.html#authentication

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RundeckAPIURL

The URL from Rundeck API. https://docs.rundeck.com/docs/api/rundeck-api.html#urls

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Required for all TPS Reports submitted after 2/19/1999

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectID

Required for all TPS Reports submitted after 2/19/1999

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ProjectId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### ProjectId

Returns the <code>ProjectId</code> value.

