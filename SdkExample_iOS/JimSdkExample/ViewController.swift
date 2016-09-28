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

    @IBOutlet weak var originLabel: UILabel!
    @IBOutlet weak var urlLabel: UILabel!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        var client: GoJimsdkClient? = nil
        GoJimsdkNewClient(&client, nil)
        
        if let client = client {
            originLabel.text = client.origin()
            urlLabel.text = client.url()
        }
    }

}

