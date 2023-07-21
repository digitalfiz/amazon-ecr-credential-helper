// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the image tag mutability settings for the specified repository. For
// more information, see Image tag mutability (https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-tag-mutability.html)
// in the Amazon Elastic Container Registry User Guide.
func (c *Client) PutImageTagMutability(ctx context.Context, params *PutImageTagMutabilityInput, optFns ...func(*Options)) (*PutImageTagMutabilityOutput, error) {
	if params == nil {
		params = &PutImageTagMutabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutImageTagMutability", params, optFns, c.addOperationPutImageTagMutabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutImageTagMutabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutImageTagMutabilityInput struct {

	// The tag mutability setting for the repository. If MUTABLE is specified, image
	// tags can be overwritten. If IMMUTABLE is specified, all image tags within the
	// repository will be immutable which will prevent them from being overwritten.
	//
	// This member is required.
	ImageTagMutability types.ImageTagMutability

	// The name of the repository in which to update the image tag mutability settings.
	//
	// This member is required.
	RepositoryName *string

	// The Amazon Web Services account ID associated with the registry that contains
	// the repository in which to update the image tag mutability settings. If you do
	// not specify a registry, the default registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type PutImageTagMutabilityOutput struct {

	// The image tag mutability setting for the repository.
	ImageTagMutability types.ImageTagMutability

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutImageTagMutabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutImageTagMutability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutImageTagMutability{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutImageTagMutabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutImageTagMutability(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutImageTagMutability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "PutImageTagMutability",
	}
}
