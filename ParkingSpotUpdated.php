<?php

namespace App\Events;

use Illuminate\Broadcasting\Channel;
use Illuminate\Broadcasting\InteractsWithSockets;
use Illuminate\Broadcasting\PrivateChannel;
use Illuminate\Contracts\Broadcasting\ShouldBroadcast;
use Illuminate\Foundation\Events\Dispatchable;
use Illuminate\Queue\SerializesModels;
use App\Models\ParkingSpot;

class ParkingSpotUpdated implements ShouldBroadcast
{
    use Dispatchable, InteractsWithSockets, SerializesModels;

    public $parkingSpot;

    public function __construct($parkingSpot)
    {
        $this->parkingSpot = $parkingSpot;
    }

    public function broadcastOn()
    {
        return new Channel('parking');
    }
}
