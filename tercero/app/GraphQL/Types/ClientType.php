<?php
namespace App\GraphQL\Types;

use App\Models\Client;
use GraphQL\Type\Definition\Type;
use Rebing\GraphQL\Support\Type as GraphQLType;

class ClientType extends GraphQLType
{
    protected $attributes = [
        'name' => 'Client',
        'description' => 'A type',
        'model' => Client::class,
    ];

    public function fields(): array
    {
        return [
            'id' => [
                'type' => Type::nonNull(Type::id()),
                'description' => 'The id of the client',
            ],
            'name' => [
                'type' => Type::string(),
                'description' => 'The name of the client',
            ],
            'email' => [
                'type' => Type::string(),
                'description' => 'The email of the client',
            ],
            'phone' => [
                'type' => Type::string(),
                'description' => 'The phone of the client',
            ],
        ];
    }
}
