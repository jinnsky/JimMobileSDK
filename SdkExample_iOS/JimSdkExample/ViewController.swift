//
//  ViewController.swift
//  JimSdkExample
//
//  Created by 杨景天 on 27/09/2016.
//  Copyright © 2016 JinnSky. All rights reserved.
//

import UIKit
import Jimsdk

class ViewController: UIViewController {
    
    @IBOutlet weak var contentLabel: UILabel!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        var client: GoJimsdkClient? = nil
        GoJimsdkNewClient("http://api2.jimyun.com", "", "", "", &client, nil)
        
        if let client = client {
            contentLabel.text = String(client.serverTimestamp())
        }
    }

}

